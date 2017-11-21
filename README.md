# MEmu5

MEmu5 is an attempt to digitalise the abstract rules for the Matrix in Shadowrun 5E.

## Get started

1. Edit "playerDB.txt" to add your character using template provided. 
2. run MEmu5.exe
3. activate your character: ```login>[charName]``` (Decker is the name of test character)
4. search host: ```matrix search>host>[hostName]```
5. do your shadowstuff using rules in Shadowrun 5 Core Rule Book (p. 217-249)


Latest Build can be found here:
https://www.dropbox.com/sh/typ2ztf1qtn9t33/AAC2CKFcq8vNDA-udDCWYwG5a?dl=0


## Supported features

### Hosts

If you search a host that is unknown to the program, it will 
create a new random host with the name you added to the
command. After that it will add the host to "HostDb.txt" so
you are able to hack it.

### IC

IC created and remembered by the program. Implemented all ICes
from Core Rule Book and Data Trails.

### Devices

Implemented the Cyberdecks from Core Rule Book

### Programs

implemented most of the programs from Core Rule Book and Data Trails

### Dice Control

you can spend Edge using "Push the Limit", "Negate Glitch" and "Reroll" interrupt actions

## What I plan to do:

* implementations of all devices from Shadowrun5 universe
* implementations of all Technomancer stuff
* implementations of all AI stuff
* A LOT of bugs fixes
* updating as new books and features come to SR Universe
* Learn to do better coding ;)

## I would be glad to receive:

Full (detailed) description of bugs you encounter

Ideas and maybe advice on how to do a better code

You can send it all to galdoba@mail.ru

## Actions Guide

To initiate an action just type:
```[ACTION_NAME]>[TARGET_NAME]>[OPTIONAL_ARGUMENTS]```
Input is NOT caSE SEnseTiVe

"Login":
```
login>decker	# will read "PlayerDB.txt" and get all info on character named "Decker"
```

"Hack on the Fly":
optional arguments: "-2m", "-3m", "-pl"
```
hack on the fly>aresmain           # will perform HotF action on host with name "Aresmain"
hack on the fly>file 17>-2m         # will perform HotF action on file with name "File 17" with modifier -4 to Dicepool in attempt to place 2 MARKs
hack on the fly>patrol ic           # will perform HotF action on IC with name "Patrol IC"
hack on the fly>TrACk ic>-3m>-pl	# will perform HotF action on IC with name "Track IC" with modifier -10 to Dicepool in attempt to place 3 MARKs using "Push the Limits" rules
```

"Brute Force":
optional arguments: "-2m", "-3m", "-pl"
```
brute force>bravo3           # will perform BF action on host with name "Bravo3"
brute force>file 17>-2m         # will perform BF action on file with name "File 17" with modifier -4 to Dicepool in attempt to place 2 MARKs
brute force>patrol ic           # will perform BF action on IC with name "Patrol IC"
brute force>TrACk ic>-3m>-pl	# will perform BF action on IC with name "Track IC" with modifier -10 to Dicepool in attempt to place 3 MARKs using "Push the Limits" rules
```

"Crack File":
optional arguments: "-pl"
```
crack file>file 9           # will perform Crack file action on file with name "File 9"
```

"Check Overwatch Score":
optional arguments: "-pl"
```
check overwatch score           # will perform inform player on current OS acording to rules in CRB. This is alternative to running "Baby Monitor" program
```

"Disarm data bomb":
optional arguments: "-pl"
```
disarm databomb>file 9           # will perform Disarm Data Bomb action on file with name "File 9"
```

"Enter Host":
```
enter host>bravo3           # will enter to host with name "Bravo3" if player has MARK on that host
```

"Exit Host":
```
exit host           # will leave current host if player is not Link-locked
```

"Erase mark":
optional arguments: "-pl"
```
erase mark           # will atempt to delete 1 random MARK placed on player (will add option to delete mark on another Icons)
```

"Edit File":
optional arguments: "copy", "delete", "download", "encrypt"
```
Edit>file 9           # will do... nothing. Will run test and change "Last Edit Date" of the file to current time. Needed mostly for RolePlaing purposes. (Not implemented properly) 
Edit>file 9>copy           # will create a new file with ownership of the player
Edit>file 9>delete           # will destroy file
Edit>file 9>download           # will initiate download process. (at the end of turn (Data Processing*5) Mp will be transfered to players device
Edit>file 9>encrypt           # will set Encryption Rating according to rules in CRB (Not implemented yet)
```

"Grid Hop":
```
grid hop>local grid           # change current grid to grid with the name "Local Grid" (so far all grids are open - will add restrictions in later versions)
```

"Load Program":
```
load program>shell           # load program "Shell". (Most programs are working and so far all available for a player. Look for descryptions in CRB and DT)
```

"Swap Attributes":
```
swap attributes>attack>sleaze           # will swap "Attack" and "Sleaze" attributes on player's device if device allows
```

"Swap Programs":
```
swap programs>armor>baby monitor           # unload program "Armor" and load program "Baby Monitor" in it's place
```

"Switch Interface Mode":
```
switch interface mode>ar       # set your mode to AR acording to rules in CRB
switch interface mode>cold       # set your mode to COLD-SIM VR acording to rules in CRB
switch interface mode>hot       # set your mode to HOT-SIM VR acording to rules in CRB
```

"Matrrix Perception":
optional arguments: "-pl", "all"
```
matrix perception>all			#Scan environment (Current Host) in attempt to find hidden (silent running) icons. Will inform how many hidden icons are there and will try to spot 1 random icon
matrix perception>file 10	# scan icon with name "file 10"
```

"Matrix search":
optional arguments: "-pl"
```
matrix search>host>alpha        # will initiate search for icon type "Host" with name "Alpha" (will take few combat turns acording to rules in Core Rule Book)
matrix search>file>system log   # will initiate search for icon type "File" with random name and filename "SYSTEM_LOG" (will take few combat turns acording to rules in Core Rule Book)
```

"Wait": 
(Nessesary after you initiated search or download)
```
wait                #wait until End of Combat Turn
wait>14             #drop players initiative by 14
wait>-ev            #wait until any "Search" or "Download" process is complete
```



## Disclaimer

All products related to Shadowrun are the property of Catalyst Game Labs.
If you interested in what this stuf is about please check their page:
https://www.catalystgamelabs.com/
