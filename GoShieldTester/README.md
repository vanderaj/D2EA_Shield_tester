# Introduction

This is the Go port of Down To Earth Astronomy's Elite Dangerous Shield Tester

D2EA's original Powershell version can be found here:
https://github.com/DownToEarthAstronomy/D2EA_Shield_tester

# Abstract

Many of us run many different ships, with many stored shield generators and 
modules with many forms of engineering. It might be tempting to just put on 
Heavy Duty / Deep plating, but is that really the best alternative? How do 
you choose the best loadout? 

Before D2EA's shield tester tool, it was the usual metas, which undeniably work. However, there's 
so many combinations, it's hard to say for sure if the meta for your ship and combat or defence 
scenario is the best alternative. 

We need a way of figuring out the best combination of generator and shield boosters for situational
scenarios. For example, you might want to change between mining to fighting Thargoid Interceptors 
or NPCs in combat zones. All of three scenarios require slightly different loadouts. 

The Go port helps you quickly find the best starting loadout for your scenario.

## Why a Go version? 

tl;dr: Speed. Nothing else. The other versions work just fine. 

The original Powershell version is groundbreaking research, but is fairly slow,
and thus might discourage some from running the tool when they change ships or
combat scenarios. 

The multi-threaded Go port is about 10-15,000 times faster per CPU thread. 
There are no loadout combinations that take more than 2 seconds on a modern
 i7, whereas the same configuration using the PowerShell version on the 
 same system can take around 10 minutes, and the Python version takes ~10 seconds.

## Improvements to these tools

In a comment to the original D2EA video, Cmdr Kaethena listed a few limitations and scenarios that you should read to understand
that these tools are a good starting point, but possibly not the ending point for your shield loadouts. There are a lot of situations where
a more generalist loadout might help you more than a max survivability loadout from this tool. YMMV. 

# Building

Building the tool is really simple:

```
go build .
```

# Running

```
go run . 
```

After building, you can also invoke it like any other command line tool. If you invoke it with the -h flag, you'll get the following:

```
  -adps float
        Absolute DPS percentage (use 100 for Thargoids) (default 200)
  -boosters int
        Number of Shield Boosters (default 2)
  -cucumber
        Useful Cucumber defaults
  -dmg float
        Damage effectiveness (use 0.1 for PvE, 0.5 for PvP, and 0.65 for Thargoids) (default 0.65)
  -edps float
        Explosive DPS percentage (use 0 for Thargoids)
  -fullboost
        Load the full booster list
  -gshp float
        Guardian HitPoints (default 0)
  -kdps float
        Kinetic DPS percentage (use 0 for Thargoids)
  -noprismatics
        Disable Prismatic shields
  -scb float
        SCB HitPoints (default 0)
  -tdps float
        Thermal DPS percentage (use 0 for Thargoids)
  -thargoid
        Useful Thargoid defaults
```

# Usage

The Go port has a number of command line flags to override the default configuration. 

## Scenario flags 

Included are two scenarios loadout configurations to aid in testing the tool, but also generate useful loadouts for the discerning
Commander who might be fighting Thargoids or D2EA. 

Disabling prismatic shields allows for users who have yet to unlock prismatics to see what their next best alternative might be. 

```
  -cucumber
        Useful Cucumber defaults
  -thargoid
        Useful Thargoid defaults
  -noprismatics
        Disable prismatic shields 
```

Those who have watched the video will recognize the defaults in use. As these override the defaults, using the other DPS flags 
doesn't work. This is a known limitation that might be fixed in a future version.

## Full booster variants

For speed's sake, the original Powershell version excluded a few booster alternatives. This is not necessary for the 
Go version, but for compatibility's sake, it is the default choice. If you want to test all boosters, use the following flag:

```
  -fullboost
        Load the full booster list
```

## Number of boosters

Choose between 0 and 8 boosters to fill up those utility slots. 

```
  -boosters int
        Number of Shield Boosters (default 2)
```

## Damage Effectiveness

Damage effectiveness is the percentage of time you'll be taking fire. Something like a PvP Commander who is using turrets might
be able to hit you say 65% of the time, so use 0.65. A Cmdr using fixed plasma or rail weapons will hit you may be 10% of the time, allowing 
you to regenerate your shields between hits. Obviously the latter really hurt, so ... really up that DPS when you lower this score.

```
  -dmg float
        Damage effectiveness (use 0.1 for PvP with fixed, 0.5 for PvE or PvP with gimballs, and 0.25 for Thargoids) (default 0.5)
```

## Damage Per Second Flags

These are expressed in percentages. A well balanced build should try to make these add up to 100%. Of course you can make these add up to 
anything, but it may affect your survivability and ability to take all forms of damage because the game and this tool may not necessarily
work out modifiers in the same way above 100%. 

```
  -adps float
        Absolute DPS percentage (use 100 for Thargoids)
  -edps float
        Explosive DPS percentage (use 0 for Thargoids) (default 33)
  -kdps float
        Kinetic DPS percentage (use 0 for Thargoids) (default 33)
  -tdps float
        Thermal DPS percentage (use 0 for Thargoids) (default 33)
```

## SCBs

If you have SCBs, include their regen hitpoints here. As you can only use one at a time, only include one. 

```
  -scb float
        SCB HitPoints (default 0)
```

## Guardian Shield Boosters

If you have Guardian Shield Boosters, include their combined hitpoints here. 

```
  -gshp float
        Guardian HitPoints (default 0)
```

# Example run

```
> .\GoShieldTester.exe -boosters 8 -fullboost -kdps 33 -tdps 33 -edps 33 -adps 0 -dmg 0.50     
Down to Earth Astronomy's ShieldTester (https://github.com/DownToEarthAstronomy/D2EA_Shield_tester)
Go port by Andrew van der Stock, vanderaj@gmail.com

Loading all boosters
Test started at:  Sun, 03 Nov 2019 17:08:40 MST
Loaded 45 generator variants
Loaded 20 shield booster variants
Loadout shield booster variations to be tested per generator:  2220075
Total loadouts to be tested:  99903375
Tests [#############################################]
Time elapsed: [ 827.7854ms ]


---- TEST SETUP ----
Shield Booster Count: [ 8 ]
SCB Hit Point Pool:   [ 0 ]
Guardian HP  Pool:    [ 0 ]
Explosive DPS:        [ 33 ]
Kinetic DPS:          [ 33 ]
Thermal DPS:          [ 33 ]
Absolute DPS:         [ 0 ]
Damage Effectiveness: [ 0.5 ]

---- TEST RESULTS ----
Survival Time:        266.07 s
Shield Generator:     Prismatic with Reinforced engineering, Hi-Cap experimental
Shield Booster 1      Resistance Augmented engineering, Thermo Block experimental
Shield Booster 2      Resistance Augmented engineering, Thermo Block experimental
Shield Booster 3      Resistance Augmented engineering, Thermo Block experimental
Shield Booster 4      Heavy Duty engineering, Super Capacitors experimental
Shield Booster 5      Heavy Duty engineering, Super Capacitors experimental
Shield Booster 6      Heavy Duty engineering, Super Capacitors experimental
Shield Booster 7      Heavy Duty engineering, Super Capacitors experimental
Shield Booster 8      Heavy Duty engineering, Super Capacitors experimental

Shield Hitpoints:     5384.2 hp
Shield Regen:         1.10 hp/s
Explosive Resistance: 72.21%
Kinetic Resistance:   66.65%
Thermal Resistance:   35.16%
```
