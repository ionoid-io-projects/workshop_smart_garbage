# workshop smart garbage
Material for raspberry pi workshop

# Wiring

<!-- ![Lighting an LED](doc/img/Irrigation.png)

|RELAY                | Raspberry pi                  |
|-------------------|---------------------------------|
|-                  | GND pin                         |
|+                  | +5v pin                         |
|Signal             | GPIO 11                         |

|Soil Moisture      | Raspberry pi                  |
|-------------------|---------------------------------|
|-                  | GND pin                         |
|+                  | +5v pin                         |
|Signal             | GPIO 21                         | -->

# Code

## 1- Power your raspberry

You can achive it with connecting it to your pc trought the Micro USB Port of the raspberry pi

![power](doc/img/1-min.jpg)

## 2- Connect to your raspberry pi
Using putty if you're on windows, Ssh if you're on a linux based os
Follow the following instruction if you dont know how to connect to raspberry pi
[Connect to raspberry pi using Putty](https://github.com/ionoid-io-projects/workshop/blob/master/doc/od-iot-raspbian-rpi-zero-windows.md#5-first-boot)

## 3- Download relay binary file

Assuming you're connected with... copy and past this command
If you're using Raspberry zero
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_smart_garbage/master/bin/arm6/garbage
```

If you're using Raspberry 3 b
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_smart_garbage/master/bin/arm7/garbage
```
## make it executable
```
chmod +x garbage
```

## 4- execute binary
```
./garbage
```

## How to stop the program
To quit or stop the program click on **Ctrl+C**
