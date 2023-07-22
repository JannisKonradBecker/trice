/* USER CODE BEGIN Header */
/**
  ******************************************************************************
  * @file    stm32f0xx_it.c
  * @brief   Interrupt Service Routines.
  ******************************************************************************
  * @attention
  *
  * <h2><center>&copy; Copyright (c) 2021 STMicroelectronics.
  * All rights reserved.</center></h2>
  *
  * This software component is licensed by ST under BSD 3-Clause license,
  * the "License"; You may not use this file except in compliance with the
  * License. You may obtain a copy of the License at:
  *                        opensource.org/licenses/BSD-3-Clause
  *
  ******************************************************************************
  */
/* USER CODE END Header */

/* Includes ------------------------------------------------------------------*/
#include "main.h"
#include "stm32f0xx_it.h"
/* Private includes ----------------------------------------------------------*/
/* USER CODE BEGIN Includes */
#include "trice.h"
/* USER CODE END Includes */

/* Private typedef -----------------------------------------------------------*/
/* USER CODE BEGIN TD */

/* USER CODE END TD */

/* Private define ------------------------------------------------------------*/
/* USER CODE BEGIN PD */

/* USER CODE END PD */

/* Private macro -------------------------------------------------------------*/
/* USER CODE BEGIN PM */

/* USER CODE END PM */

/* Private variables ---------------------------------------------------------*/
/* USER CODE BEGIN PV */

// us64 is a 64-bit microsecond counter, counting circular in steps of 1000 every ms.
static uint64_t us64 = 0;

// us16 is a 16-bit microsecond counter, running parallel to us64, but is reset every 10ms, so gets these values: 0, 1000, ... 9000.
static uint16_t us16 = 0;

// ms32 is a 32-bit millisecond counter, counting circular in steps of 1 every ms.
static uint32_t ms32 = 0;

// ms16 is a 16-bit millisecond counter, running parallel to ms32, but is reset every 10s.
static uint16_t ms16 = 0;

// us64_1 holds the result of last Us16() or Us64() call.
static uint64_t us_1 = 0;

/* USER CODE END PV */

/* Private function prototypes -----------------------------------------------*/
/* USER CODE BEGIN PFP */

/* USER CODE END PFP */

/* Private user code ---------------------------------------------------------*/
/* USER CODE BEGIN 0 */

//! US_DUTY is a helper macro to avoid code duplication.
//! US_DUTY updates the us_1 tick in the assumption of an 48MHz systick clock using the current systick value.
//! ATTENTION: This is a quick and dirty implementation working well only if this function is called in intervals smaller than 1 ms.
//! :-( Because the STM32F030 has no 32-bit sysclock counter we need to compute this value or concatenate two 16-bit timers. )
//! I see no way to find out if the systick ISR was already active shortly after a systick counter wrap, despite calling this
//! functionality in intervals smaller than 1 ms if not using hardware timers. To make it clear: You can use Us64() to measure long
//! intervals up to 584542 years, but the "OS" needs to call Us64 internally regularely in <1ms intervals.
#define US_DUTY \
    uint32_t usOffset = (((SysTick->LOAD - SysTick->VAL) * 87381) >> 22); /* Divide 48MHz clock by 48,0001831 to get us part. */ \
    uint64_t us = us64 + usOffset;                                        /* 47999*87381 < 2^32 */ \
    int correction = 0; \
    if( us < us_1){ /* Possible very close to systick ISR, when us64 was not incremented yet, but the systic wrapped already. */ \
        correction = 1000; /* Time cannot go backwards, so correct the 1ms error in the assumption last call is not longer than 1ms back. */ \
    } \
    us_1 = us + correction; /* keep for next call */

//  //! Us64 reads the 1us tick.
//  //! \retval us count since last reset
//  static inline uint64_t Us64( void ){
//      US_DUTY
//      return us_1;
//  }

//! Us32 reads the 1us tick. A 32-bit us stamp, wraps after 71,58 seconds.
//! \retval us count since last reset
uint32_t Us32( void ){
    US_DUTY
    return us_1;
}

//  //! Us16 or Us64 needs be called in < 1ms intervals.
//  static inline uint16_t Us16( void ){
//      US_DUTY
//      return us16 + usOffset + correction; // max 9000 + max 999 + max 1000 = max 10999 us;
//  }

// UsDuty is usable in wait loops, waiting for a hardware flag for example.
void UsDuty( void ){
    { US_DUTY }
    __WFE(); // wait for event (sleep)
    { US_DUTY } 
}

// usDuty is usable in short wait loops, waiting for a hardware flag for example.
void usDuty( void ){
    US_DUTY
}

uint32_t milliSecond( void ){
    return ms32;
}

/* USER CODE END 0 */

/* External variables --------------------------------------------------------*/

/* USER CODE BEGIN EV */

/* USER CODE END EV */

/******************************************************************************/
/*           Cortex-M0 Processor Interruption and Exception Handlers          */
/******************************************************************************/
/**
  * @brief This function handles Non maskable interrupt.
  */
void NMI_Handler(void)
{
  /* USER CODE BEGIN NonMaskableInt_IRQn 0 */
    TRice( iD(6409), "CRITICAL:NMI!\n" );
  /* USER CODE END NonMaskableInt_IRQn 0 */
  /* USER CODE BEGIN NonMaskableInt_IRQn 1 */
  while (1)
  {
  }
  /* USER CODE END NonMaskableInt_IRQn 1 */
}

/**
  * @brief This function handles Hard fault interrupt.
  */
void HardFault_Handler(void)
{
  /* USER CODE BEGIN HardFault_IRQn 0 */
    TRice( iD(1546), "FATAL:HardFault!\n" );
  /* USER CODE END HardFault_IRQn 0 */
  while (1)
  {
    /* USER CODE BEGIN W1_HardFault_IRQn 0 */
    /* USER CODE END W1_HardFault_IRQn 0 */
  }
}

/**
  * @brief This function handles System service call via SWI instruction.
  */
void SVC_Handler(void)
{
  /* USER CODE BEGIN SVC_IRQn 0 */

  /* USER CODE END SVC_IRQn 0 */
  /* USER CODE BEGIN SVC_IRQn 1 */

  /* USER CODE END SVC_IRQn 1 */
}

/**
  * @brief This function handles Pendable request for system service.
  */
void PendSV_Handler(void)
{
  /* USER CODE BEGIN PendSV_IRQn 0 */

  /* USER CODE END PendSV_IRQn 0 */
  /* USER CODE BEGIN PendSV_IRQn 1 */

  /* USER CODE END PendSV_IRQn 1 */
}

/**
  * @brief This function handles System tick timer.
  */
void SysTick_Handler(void)
{
  /* USER CODE BEGIN SysTick_IRQn 0 */
    us64 += 1000;
    us16 += 1000;
    us16 = us16 < 10000 ? us16 : 0;
    ms32++;
    ms16++;
    ms16 = ms16 < 10000 ? ms16 : 0; 
  /* USER CODE END SysTick_IRQn 0 */

  /* USER CODE BEGIN SysTick_IRQn 1 */
#if defined( TRICE_UARTA ) && ((TRICE_BUFFER == TRICE_DOUBLE_BUFFER) || (TRICE_BUFFER == TRICE_RING_BUFFER) ) // buffered out to UART
    triceTriggerTransmitUartA();
#endif
#if defined( TRICE_UARTB ) && ((TRICE_BUFFER == TRICE_DOUBLE_BUFFER) || (TRICE_BUFFER == TRICE_RING_BUFFER) ) // buffered out to UART
    triceTriggerTransmitUartB();
#endif
  /* USER CODE END SysTick_IRQn 1 */
}

/******************************************************************************/
/* STM32F0xx Peripheral Interrupt Handlers                                    */
/* Add here the Interrupt Handlers for the used peripherals.                  */
/* For the available peripheral interrupt handler names,                      */
/* please refer to the startup file (startup_stm32f0xx.s).                    */
/******************************************************************************/

/**
  * @brief This function handles USART2 global interrupt.
  */
void USART2_IRQHandler(void)
{
  /* USER CODE BEGIN USART2_IRQn 0 */
#if defined( TRICE_UARTA )
    if (LL_USART_IsActiveFlag_RXNE(TRICE_UARTA) ) { // Read Data Register Not Empty Flag 
     // todo: Do it better:
     // https://github.com/Apress/Beg-STM32-Devel-FreeRTOS-libopencm3-GCC/rtos/libwwg/src/getline.c
     // And put space separated args in a string array.
        static char rxBuf[TRICE_COMMAND_SIZE_MAX+1]; // with terminating 0
        static int index = 0;
        uint8_t v;
        if( LL_USART_IsActiveFlag_ORE(TRICE_UARTA) ){
            TRICE( Id(4902), "WARNING:USARTq OverRun Error Flag is set!\n" );
        }
        v = LL_USART_ReceiveData8(TRICE_UARTA); // implicit clears the flag
        rxBuf[index] = (char)v;
        index += index < TRICE_COMMAND_SIZE_MAX ? 1 : 0; 
        if( v == 0 ){ // command end
            TRICE_S( Id(5682), "rx:received command:%s\n", rxBuf );
            strcpy(triceCommandBuffer, rxBuf );
            triceCommandFlag = 1;
            index = 0;
        }
        return;
    }
#endif // #if defined( TRICE_UARTA )
    // If both flags active and only one was served, the IRQHandler gets activated again.

#if defined( TRICE_UARTA ) && ((TRICE_BUFFER == TRICE_DOUBLE_BUFFER) || (TRICE_BUFFER == TRICE_RING_BUFFER) ) // buffered out to UARTA
    if( LL_USART_IsActiveFlag_TXE(TRICE_UARTA) ){ // Transmit Data Register Empty Flag
        triceServeTransmitUartA();
        return;
    }
#endif

    /*
    if( LL_USART_IsActiveFlag_PE(TRICE_UARTA) ){
        TRICE( Id(4863), "err:TRICE_UARTA Parity Error Flag is set\n" );
        LL_USART_ClearFlag_PE(TRICE_UARTA);
    }
    if( LL_USART_IsActiveFlag_FE(TRICE_UARTA) ){
        TRICE( Id(6299), "err:TRICE_UARTA Framing Error Flag is set\n" );
        LL_USART_ClearFlag_FE(TRICE_UARTA);
    }
    if( LL_USART_IsActiveFlag_NE(TRICE_UARTA) ){
        TRICE( Id(3239), "err:TRICE_UARTA Noise error detected Flag is set\n" );
        LL_USART_ClearFlag_NE(TRICE_UARTA);
    }
    if( LL_USART_IsActiveFlag_ORE(TRICE_UARTA) ){
        //TRICE( Id(5760), "err:TRICE_UARTA OverRun Error Flag is set\n" );
        LL_USART_ClearFlag_ORE(TRICE_UARTA);
    }
    if( LL_USART_IsActiveFlag_IDLE(TRICE_UARTA) ){
        TRICE( Id(5244), "wrn:TRICE_UARTA IDLE line detected Flag is set\n" );
        LL_USART_ClearFlag_IDLE(TRICE_UARTA);
    }
    if( LL_USART_IsActiveFlag_TC(TRICE_UARTA) ){
        //TRICE0( Id(5084), "wr:TRICE_UARTA Transmission Complete Flag is set\n" );
        LL_USART_ClearFlag_TC(TRICE_UARTA);
    }
    if( LL_USART_IsActiveFlag_nCTS(TRICE_UARTA) ){
        TRICE( Id(7871), "wrn:TRICE_UARTA CTS interrupt Flag is set\n" );
        LL_USART_ClearFlag_nCTS(TRICE_UARTA);
    }
    if( LL_USART_IsActiveFlag_CTS(TRICE_UARTA) ){
        TRICE( Id(6491), "wrn:TRICE_UARTA CTS Flag is set\n" );
    }
    if( LL_USART_IsActiveFlag_RTO(TRICE_UARTA) ){
        TRICE( Id(4276), "err:TRICE_UARTA Receiver Time Out Flag is set\n" );
        LL_USART_ClearFlag_RTO(TRICE_UARTA);
    }
    if( LL_USART_IsActiveFlag_ABRE(TRICE_UARTA) ){
        TRICE0( Id(2465), "err:TRICE_UARTA Auto-Baud Rate Error Flag is set\n" );
    }
    if( LL_USART_IsActiveFlag_ABR(TRICE_UARTA) ){
        TRICE( Id(5533), "wrn:TRICE_UARTA Auto-Baud Rate Flag is set\n" );
    }
    if( LL_USART_IsActiveFlag_BUSY(TRICE_UARTA) ){
        TRICE( Id(3815), "inf:TRICE_UARTA Busy Flag is set\n" );
    }
    if( LL_USART_IsActiveFlag_CM(TRICE_UARTA) ){
        TRICE( Id(2363), "wrn:TRICE_UARTA Character Match Flag is set\n" );
        LL_USART_ClearFlag_CM(TRICE_UARTA);
    }
    if( LL_USART_IsActiveFlag_SBK(TRICE_UARTA) ){
        TRICE( Id(7804), "wrn:TRICE_UARTA Send Break Flag is set\n" );
    }
    if( LL_USART_IsActiveFlag_RWU(TRICE_UARTA) ){
        TRICE( Id(5010), "wrn:TRICE_UARTA Receive Wake Up from mute mode Flag is set\n" );
    }
    if( LL_USART_IsActiveFlag_TEACK(TRICE_UARTA) ){
        //TRICE( Id(4009), "wr:TRICE_UARTA Transmit Enable Acknowledge Flag is set\n" );
    }
    if( LL_USART_IsActiveFlag_REACK(TRICE_UARTA) ){
        //TRICE( Id(2764), "rd:TRICE_UARTA Receive Enable Acknowledge Flag is set\n" );
    }
*/
  /* USER CODE END USART2_IRQn 0 */
  /* USER CODE BEGIN USART2_IRQn 1 */

  /* USER CODE END USART2_IRQn 1 */
}

/* USER CODE BEGIN 1 */

/* USER CODE END 1 */
