# MEmu5

MEmu5 is an attempt to digitalise the abstract rules for the Matrix in Shadowrun 5E.

## Get started

1. Edit "playerDB.txt" to add your character using template provided.
2. run Sr5MEmu.exe
3. activate your character: ```login>[charName]```
4. search host: ```matrix search>host>[hostName]```
5. do your shadowstuff using rules in Shadowrun 5 Core Rule Book (p. 217-249)

## Actions Examples

to initiate an action just type:
```[ACTION_NAME]>[TARGET_NAME]>[OPTIONAL_ARGUMENTS]```

load character/his characteristics and abilities with name "Testpc"
```
login>Testpc
```

Do "Hack on the Fly" action on icon with name "Ares Host", 
optional arg "-2m" stands for 2 marks (put 2 marks with modifier -4 to Dice Pool)
```
hack on the fly>Ares Host>-2m	

hack on the fly>file 17>-2m
hack on the fly>patrol ic
hack on the fly>TrACk ic>-3m>-pl	# "-pl" is an optional argument stand for "Push the Limit"
```

Set your mode to COLD-SIM VR acording to rules in CRB
```
switch interface mode>cold vr 
```

Scan all environment in attempt to find hidden (silent running) icons.
Will inform you how many hidden icons are in the same enviroment as 
the player and will spot 1 random icon
```
matrix perception>all			
matrix perception>file 10	# scan icon with name "file 10"
```

Will initiate search for icon type "Host" with name "Alpha" 
(will take few combat turns acording to rules in Core Rule Book)
```
matrix search>host>alpha
matrix search>file>system log
```

Set players initiative to 0, and start new Combat Turn (Nessesary after you initiated search or download)
```
wait
```

And those work according the rules from the book:
```
crack file
data spike
brute force
```

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

full (detailed) description of bugs you encounter

Ideas and maybe advice on how to do a better code

## Disclaimer

All products related to Shadowrun are the property of Catalyst Game Labs.
If you interested in what this stuf is about please check their page:
https://www.catalystgamelabs.com/
