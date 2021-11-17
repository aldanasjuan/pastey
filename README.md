# Pastey

Little program to paste information with custom hotkeys. For example, ctrl+alt+1 would paste "hello" while ctrl+alt+2 would paste "bye".

## Install

- Add pastey.exe to your PATH so you can run it from a terminal.
- Create a config file with your hotkeys and data to paste using the proper sintax. (View keys.config for an example)
- Run `pastey C:/path/to/file.config`

## Sintax

- Add your modifiers (alt, ctrl, shift and win), joined with a '+' sign. 
- After the modifiers the desired key like "1", "backspace", etc. (view the the file "key-list.txt" for the full list of keys). For numbers enter "1" instead of "one", the program will convert it properly. 
- Add a '=' sign
- Add your value to paste

Like so: 

```
ctrl+alt+1=hello
ctrl+alt+2=world
ctrl+alt+3=goodbye
```

## Reserved hotkeys

- alt+backspace (exits the program)

## TODO:
[] - Reload config file
