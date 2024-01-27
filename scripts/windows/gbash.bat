@ECHO OFF

REM Logging to Git Bash Termianl

PUSHD C:\Program Files\Git\bin
bash.exe -i -l
POPD
SET root=C:\Users\Admin\GitHub
CD %root%
