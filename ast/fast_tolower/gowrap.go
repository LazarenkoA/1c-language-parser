package fast_tolower

/*
#cgo CFLAGS: -std=c99
#include <stdlib.h>  // for free
#include <ctype.h>   // for tolower
#include <stdint.h>
#include <stdio.h> // for printf

void fastToLower(char* str) {
     for (int i = 0; str[i] != '\0';) {
        unsigned char byte = str[i];
		int currentChar = 0;

        if (byte <= 0x7F) {
            currentChar = byte; // Один байт (ASCII)
            i++;
        } else if (byte >= 0xC0 && byte <= 0xDF) {
            currentChar = (str[i] & 0x1F) << 6 | (str[i+1] & 0x3F);
            i += 2;
        } else if (byte >= 0xE0 && byte <= 0xEF) {
            currentChar = (str[i] & 0x0F) << 12 | (str[i+1] & 0x3F) << 6 | (str[i+2] & 0x3F);
            i += 3;
        } else if (byte >= 0xF0 && byte <= 0xF7) {
            currentChar = (str[i] & 0x07) << 18 | (str[i+1] & 0x3F) << 12 | (str[i+2] & 0x3F) << 6 | (str[i+3] & 0x3F);
            i += 4;
        }

		if (currentChar >= 'A' && currentChar <= 'Z') {
			str[i-1] = tolower(str[i]);
		} else if (currentChar >= 1040 && currentChar <= 1071) { // диапазон А - Я
			int codepoint = currentChar + 32;
	 		str[i-2] = 0xC0 | (codepoint >> 6);   // Старший байт
			str[i-1] = 0x80 | (codepoint & 0x3F); // Младший байт
		} else if (currentChar == 1025) { // Ё
			str[i-2] = 209;
			str[i-1] = 145;
		}
    }

}

*/
import "C"
import "unsafe"

// Go-обертка для вызова C функции
func FastToLower(str string) string {
	cs := C.CString(str)

	C.fastToLower(cs)
	result := C.GoString(cs)
	C.free(unsafe.Pointer(cs))

	return result // *(*string)(unsafe.Pointer(&runes))
}

/*
void fastToLower(char* str) {
     for (int i = 0; str[i] != '\0';) {
        unsigned char byte = str[i];
        if (byte <= 0x7F) {
            printf("%d ", byte); // Один байт (ASCII)
            i++;
        } else if (byte >= 0xC0 && byte <= 0xDF) {
            printf("%d ", (str[i] & 0x1F) << 6 | (str[i+1] & 0x3F));
            i += 2;
        } else if (byte >= 0xE0 && byte <= 0xEF) {
            printf("%d ", (str[i] & 0x0F) << 12 | (str[i+1] & 0x3F) << 6 | (str[i+2] & 0x3F));
            i += 3;
        } else if (byte >= 0xF0 && byte <= 0xF7) {
            printf("%d ", (str[i] & 0x07) << 18 | (str[i+1] & 0x3F) << 12 | (str[i+2] & 0x3F) << 6 | (str[i+3] & 0x3F));
            i += 4;
        }
    }
    printf("\n");
}


// передача рун
#include <stdlib.h>  // fFor free
#include <ctype.h>   // for tolower
#include <stdint.h>
// #include <stdio.h> // for printf

void fastToLower(int32_t* runes, int length) {
    for (int i = 0; i < length; i++) {
		if (runes[i] >= 'A' && runes[i] <= 'Z') {
			runes[i] = tolower(runes[i]);
		} else if (runes[i] >= 1040 && runes[i] <= 1071) { // диапазон А - Я
			runes[i] += 32;
		} else if (runes[i] == 1025) { // Ё
			runes[i] = 1105;
		}
    }
}

*/
