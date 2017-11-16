Ok, what this all about:

1. This is an attempt to digitalise hell of an abstract rules for matrix Chapter in Shadowrun 5E

2 HOW TO USE:
-Edit "playerDB.txt" to add your character using template provided.
-run Sr5MEmu.exe:
--to activate your character type "login>[charName]" command
--search host using "matrix search>host>[hostName]" command
--do your shadowstuf using rules in Shadowrun 5 Core Rule Book (p. 217-249)

to initiate an action just type:
[ACTION_NAME]>[TARGET_NAME]>[OPTIONAL_ARGUMENTS]
Actions Examples:
login>TestPC - load character/his characteristics and abilities with name "TestPC"
"hack on the fly>Ares Host>-2m" - do "HotF" action on icon with name "Ares Host", optional arg "-2m" stands for 2 marks (put 2 marks with modifier -4 to Dice Pool)
-valid examples:
"hack on the fly>file 17>-2m"
"hack on the fly>patrol ic"
"hack on the fly>TrACk ic>-3m>-pl" - "-pl" is an optional argument stand for "Push the Limit"
"switch interface mode>cold vr" - set your mode to COLD-SIM VR acording to rules in CRB
"matrix perception>all" - scan all envirement in attempt to find hidden (silent running) icons - will inform on how many hidden icons in the same enviroment as player an will spot 1 random icon
"matrix perception>file 10" - scan icon with name "file 10"
"matrix search>host>alpha" - will initiate search for icon type "Host" with name "Alpha" (will take few combat turns acording to rules in CRB)
-valid examples:
"matrix search>file>system log"
"wait" - set players initiative to 0, and start new Combat Turn (Nessesary after you initiated search or download)
crack file/data spike/brute force - work acording the rules from the book


3. suppoted features:
-Hosts
If you search host that is unknown to program - it will create new random host with name you added to command, after that it will add the host to "HostDb.txt" so you will be able to hack exact same Host
-IC
IC created and remembered by the program.
Implemented all ICes from CRB and DT books
-devices
implemented Cyberdecks from CRB
-programs
implemented most of the programs from CRB and DT books
-dice control:
you can spend Edge using "Push the Limit", "Negate Glitch" and "Reroll" interrupt actions

4 - what I plan to do:
implementations of all devices from Shadowrun5 universe
implementations of all Technomancer stuff
implementations of all AI stuff
A LOT of bugs fixes
updating as new books and features come to SR Universe
Learn to do better codding)))

5 - I would be glad to receive:
full (detailed) description of bugs you encounter
Idead and mabe advice on how to do a better code



///////////////////////////////////

If you interested source code can be found here:
https://github.com/Galdoba/MEmu5

///////////////////////////////////

All products related to Shadowrun is the property of Catalyst Game Labs. If you interested in what this stuf is about please check theit page:
https://www.catalystgamelabs.com/

////////////////////////////////////