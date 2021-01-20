/*! \file triceConfig.h
\author Thomas.Hoehenleitner [at] seerose.net
*******************************************************************************/

#ifndef TRICE_CONFIG_H_
#define TRICE_CONFIG_H_

#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>
// Enabling next line leads to XTEA encryption  with the key. Only wrapped barel over UART us encrypted right now.
//#define ENCRYPT XTEA_KEY( ea, bb, ec, 6f, 31, 80, 4e, b9, 68, e2, fa, ea, ae, f1, 50, 54 ); //!< -password MySecret

#define TRICE_HEADLINE \
TRICE0( Id(54823), "s:                                                   \ns:   MDK-ARM_LL_UART_RTT0_BARE_STM32F030_NUCLEO-64   \ns:                                                   \n\n")

///////////////////////////////////////////////////////////////////////////////
//
//#define TRICE_ENTER_CRITICAL_SECTION { // Uncomment for more speed but only if TRICE macros & fifo access cannot
//#define TRICE_LEAVE_CRITICAL_SECTION } // get interrupted by other TRICE macros or fifo access (e.g. interrupts).
//
///////////////////////////////////////////////////////////////////////////////

#define TRICE_FIFO_BYTE_SIZE 2048 //!< must be a power of 2
#define TRICE_COMPILE TRICE_SPEED_OVER_SPACE // TRICE_SPACE_OVER_SPEED or TRICE_SPEED_OVER_SPACE

///////////////////////////////////////////////////////////////////////////////
//

//! Select target trice transfer encoding. 
#define TRICE_ENCODING TRICE_BARE_ENCODING

//! Set endianess according to target hardware. Options: TRICE_BIG_ENDIANESS, TRICE_LITTLE_ENDIANESS.
//! Some compiler offer an automatic detection for this.
#define TRICE_HARDWARE_ENDIANESS TRICE_LITTLE_ENDIANESS 

//! Set byte order according desired transfer format. Options: TRICE_BIG_ENDIANESS, TRICE_LITTLE_ENDIANESS. 
//! TRICE_BIG_ENDIANESS is network order.
//! If TRICE_TRANSFER_ENDIANESS is equal to TRICE_HARDWARE_ENDIANESS the trice code is smaller and more efficient.
//! When set to TRICE_LITTLE_ENDIANESS the trice tool -enc format specifier is extended by a letter 'L'.
//! Example -enc "pack2L" -> -enc "pack2L".
#define TRICE_TRANSFER_ENDIANESS TRICE_LITTLE_ENDIANESS 

//#define TRICE_U8PUSH(v) do{ triceU8PushSeggerRTT(v); triceU8Push(v); } while(0) //!< Set trice out channel(s) 
#define TRICE_U32PUSH(v) do{ triceU32PushSeggerRTT(v); triceU32Push(v); } while(0) //!< Set trice out channel(s) 

//
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// uncomment for trice wrap transfer format
// #define TRICE_WRAP_START_BYTE 0xEB // 235, to do: also as trice tool parameter
// #define TRICE_WRAP_LOCAL_ADDR 0x80 // to do: also as trice tool parameter
// #define TRICE_WRAP_DEST_ADDR  0x81 // to do: also as trice tool parameter
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// uncomment for trice UART transfer format
#define TRICE_UART USART2 // set UART number
///////////////////////////////////////////////////////////////////////////////

void triceCheckSet( int index ); //!< tests

#ifdef __cplusplus
}
#endif

#endif /* TRICE_CONFIG_H_ */
