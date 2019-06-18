# tankpreise

Query prices of gas

Developed using go-modules

## Library use

livingit.de/code/tankpreise is a library wrapper for https://creativecommons.tankerkoenig.de/

Example usage

    gp, err := tankpreise.NewGasPrices()
    gp.SetLicense("get one")
    p, err := gp.PriceQuery(tankpreise.PricesRequest{
        IDs: stations,
    })

stations is a string array of station ids, the license number presented here is the demo license from the provided and will return only example data

## CLI

livingit.de/code/tankpreise/cmd/tankpreise is a command line utility to query the api.

### Search for gas stations

    tankpreise search --latitude 52.521 --longitude 13.438 --radius 10 --sort dist --gas-type all 

will return gas stations in a radius with 10 around provided locations sorted by distance regardless of gas type:

    ID                                     Open    Name                                      Street                           ZipCode   City
    474e5046-deaf-4f9b-9a32-9797b778f047   true    TOTAL BERLIN                              MARGARETE-SOMMER-STR. 2          10407     BERLIN
    4429a7d9-fb2d-4c29-8cfe-2ca90323f9f8   true    TOTAL BERLIN                              HOLZMARKTSTR. 36-42              10243     BERLIN
    278130b1-e062-4a0f-80cc-19e486b4c024   true    Aral Tankstelle                           Holzmarktstraße 12/14            10179     Berlin
    1c4f126b-1f3c-4b38-9692-05c400ea8e61   true    Sprint Berlin Kniprodestr.                Kniprodestr. 25                  10407     Berlin
    ...

ID must be used in calls that provide details or price comparisons

### Price comparison

    tankpreise price --station-id 474e5046-deaf-4f9b-9a32-9797b778f047 --station-id 4429a7d9-fb2d-4c29-8cfe-2ca90323f9f8

will return list of prices for provided stations

    ID                                     Status   E5         E10        Diesel
    474e5046-deaf-4f9b-9a32-9797b778f047   open     1.234000   1.234000   1.234000
    4429a7d9-fb2d-4c29-8cfe-2ca90323f9f8   open     1.234000   1.234000   1.234000

### Details on station

    tankpreise detail --station-id 474e5046-deaf-4f9b-9a32-9797b778f047

will return details on the station

    ID              474e5046-deaf-4f9b-9a32-9797b778f047
    Name            TOTAL BERLIN
    Street          MARGARETE-SOMMER-STR. 2
    City            10407 BERLIN
    State
    Brand           TOTAL
    Opening times
    Is open         true

## History

|Version|Description|
|---|---|
|0.1.0|initial version|
