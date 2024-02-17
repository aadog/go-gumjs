**install:**

**download gum-devkit:**[click here](https://github.com/frida/frida/releases)

- windows x64:

    1.copy frida-gumjs.h to msys2 `C:\msys64\mingw64\include\frida-gumjs.h`

    2.copy libfrida-gumjs.lib to msys2 `C:\msys64\mingw64\lib\libfrida-gumjs.lib`

    3.convert msvc .lib to gcc .a

- linux x64:

  1.copy frida-gumjs.h to msys2 `/usr/include/frida-gumjs.h`

  2.copy libfrida-gumjs.lib to msys2 `/usr/lib/libfrida-gumjs.a`

- linux x32:

  1.copy frida-gumjs.h to msys2 `/usr/include/frida-gumjs.h`

  2.copy libfrida-gumjs.lib to msys2 `/usr/lib32/libfrida-gumjs.a`

    
- android arm64

  ```
  @set NDK=%ANDROID_HOME%\ndk-bundle
  @set TOOLCHAIN=%NDK%\toolchains\llvm\prebuilt\windows-x86_64
  @set SYSROOT=%TOOLCHAIN%/sysroot
  ```

  1.copy frida-gumjs.h to msys2 `%SYSROOT%/usr/include/frida-gumjs.h`

  2.copy libfrida-gumjs.lib to msys2 `%SYSROOT%/usr/lib/aarch64-linux-android/libfrida-gumjs.a`

