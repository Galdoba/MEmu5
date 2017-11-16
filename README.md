# MEmu5

MEmu5 is an attempt to digitalise the abstract rules for the Matrix in Shadowrun 5E.

## HOW TO USE:

1. Edit "playerDB.txt" to add your character using template provided.
2. run Sr5MEmu.exe
3. activate your character: ```login>[charName]```
4. search host: ```matrix search>host>[hostName]```
5. do your shadowstuff using rules in Shadowrun 5 Core Rule Book (p. 217-249)

## Actions Examples:

to initiate an action just type:
```[ACTION_NAME]>[TARGET_NAME]>[OPTIONAL_ARGUMENTS]```

load character/his characteristics and abilities with name "TestPC"
```
login>TestPC 
```

do "HotF" action on icon with name "Ares Host", 
optional arg "-2m" stands for 2 marks (put 2 marks with modifier -4 to Dice Pool)
```
hack on the fly>Ares Host>-2m	

hack on the fly>file 17>-2m
hack on the fly>patrol ic
hack on the fly>TrACk ic>-3m>-pl	# "-pl" is an optional argument stand for "Push the Limit"
```

set your mode to COLD-SIM VR acording to rules in CRB
```
switch interface mode>cold vr 
```

scan all envirement in attempt to find hidden (silent running) icons.
will inform you how many hidden icons in the same enviroment as player,
and will spot 1 random icon
```
matrix perception>all			
matrix perception>file 10	# scan icon with name "file 10"
```

will initiate search for icon type "Host" with name "Alpha" 
(will take few combat turns acording to rules in Core Rule Book)
```
matrix search>host>alpha
matrix search>file>system log
```

set players initiative to 0, and start new Combat Turn (Nessesary after you initiated search or download)
```
wait
```

And those work acording the rules from the book:
```
crack file
data spike
brute force
```

## supported features:

### Hosts

If you search host that is unknown to program, it will create
new random host with name you added to command, after that it
will add the host to "HostDb.txt" so you will be able to hack
the exact same Host

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

full (detailed) description of bugs you encounter

Ideas and maybe advice on how to do a better code

## Disclaimer

All products related to Shadowrun are the property of Catalyst Game Labs.
If you interested in what this stuf is about please check their page:
https://www.catalystgamelabs.com/
