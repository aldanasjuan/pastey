# Pastey

Little program to paste information with custom hotkeys. For example, ctrl+alt+1 would paste "hello" while ctrl+alt+2 would paste "bye".

## Install

- Add pastey.exe to your PATH so you can run it from a terminal.
- Create a config file with your hotkeys and data to paste using the proper sintax. (View keys.config for an example)
- Run `pastey C:/path/to/file.config`

## Sintax

- Add your modifiers (only "alt", "ctrl", "shift" and "win" allowed), joined with a '+' sign. 
- After the modifiers, add your desired key. Like "1", "backspace", etc. (view the the file "key-list.txt" for the full list of keys). For numbers enter "1" instead of "one", the program will convert it properly. 
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

## Notes
- You can't register the same hotkey twice. 
- If you register existing hotkeys, they will be overridden while the program runs. For example, ctrl+v would not work as regular paste, so be careful. 

## TODO:
- [ ] Reload config file
