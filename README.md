Raadhuis Xmas Tree
==================

Web front end application thingy to control the Xmas tree (hence the name) RGB
LEDs in Lieshout, the Netherlands.

What?
-----

This thing is going to do three things:

* Serve a web frontend to choose a color for the LEDs
* Parse RGB API requests and send them to the arduino below the tree
* Somehow serve the video feed from the webcam

Maybe a fourth thing:

* Being awesome

Hopefully not the fifth thing:

* Crashing

Testing
-------

Since we don't lack confidence, there are no tests.

This is the command that fakes the arduino:

    nc -k -l 9000 | ruby -e 'loop { STDIN.read(4).each_byte { |c| printf "%#X ", c.ord }; puts }'

***

           *             ,
                       _/^\_
                      <     >
     *                 /.-.\         *
              *        `/&\`                   *
                      ,@.*;@,
                     /_o.I %_\    *
        *           (`'--:o(_@;
                   /`;--.,__ `')             *
                  ;@`o % O,*`'`&\
            *    (`'--)_@ ;o %'()\      *
                 /`;--._`''--._O'@;
                /&*,()~o`;-.,_ `""`)
     *          /`,@ ;+& () o*`;-';\
               (`""--.,_0 +% @' &()\
               /-.,_    ``''--....-'`)  *
          *    /@%;o`:;'--,.__   __.'\
              ;*,&(); @ % &^;~`"`o;@();         *
              /(); o^~; & ().o@*&`;&%O\
        jgs   `"="==""==,,,.,="=="==="`
           __.----.(\-''#####---...___...-----._
         '`         \)_`"""""`
                 .--' ')
               o(  )_-\
                 `"""` `
