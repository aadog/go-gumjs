#include "go-gum.h"
#include <stdio.h>

gboolean GumCallBack(const GumModuleDetails * details,void* user_data){
    printf("hello %s\n",details->name);
    return true;
}
void testC(){
    gum_init_embedded();
    gum_process_enumerate_modules((GumFoundModuleFunc)(GumCallBack), 0);
}
