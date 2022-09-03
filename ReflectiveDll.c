#include <windows.h>
#include <stdio.h>
#include "main.h"


BOOL WINAPI DllMain(
    HINSTANCE hinstDLL,
    DWORD fdwReason,
    LPVOID lpReserved )
{
    switch( fdwReason )
    {
        case DLL_PROCESS_ATTACH:
;
            break;

        case DLL_THREAD_ATTACH:

            break;

        case DLL_THREAD_DETACH:
            break;

        case DLL_PROCESS_DETACH:
            break;
    }
    return TRUE;
}

int RunDll(){
    MyFunction();
    return 0;
}