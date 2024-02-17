**install:**

**download gum-devkit:**[click here](https://github.com/frida/frida/releases)

- windows x64:

    1.copy frida-gum.h to msys2 `C:\msys64\mingw64\include\frida-gum.h`

    2.copy libfrida-gum.lib to msys2 `C:\msys64\mingw64\lib\libfrida-gum.lib`

    3.convert msvc .lib to gcc .a

- linux x64:

  1.copy frida-gum.h to msys2 `/usr/include/frida-gum.h`

  2.copy libfrida-gum.lib to msys2 `/usr/lib/libfrida-gum.a`

  3.convert msvc .lib to gcc .a

- linux x32:

  1.copy frida-gum.h to msys2 `/usr/include/frida-gum.h`

  2.copy libfrida-gum.lib to msys2 `/usr/lib32/libfrida-gum.a`

  3.convert msvc .lib to gcc .a
    
