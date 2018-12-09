// version.c
#include <stdint.h>
#include <string.h>
#include <stdio.h>
/*
* __DATE__ : May 12 2017  月日年
* __TIME__ : 15:26:26     时分秒
* return : 2018-12-09 15:26:26
*/
static const char *months[] = {"Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"};
static char buildtime[20];

const char *Buildtime(void)
{
	uint8_t i,lengh;
	char *pstr = __DATE__;
	char *year = &pstr[strlen(pstr) - 4];
	uint8_t month;
	uint8_t day;
	
	for(i = 0;i < sizeof(months)/sizeof(months[0]); i++){
        if(memcmp(months[i],pstr,3) == 0){
            month = i + 1;
            break;
        }
    }
	
	day = pstr[4] == ' ' ? 0 : (pstr[4] - 0x30);
	day = day * 10 + pstr[5] - 0x30;
	
	(void)sprintf(buildtime,"%s-%02d-%02d %s", year, month, day, __TIME__);
	
    return buildtime;
}