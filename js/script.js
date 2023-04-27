// Copyright (c) 2022 Mariam Kasim All rights reserved
//
// Created by: Mariam Kasim
// Created on: Apr 26
// This file contains the JS functions for index.html

"use strict"

const randomNumber = Math.floor(Math.random() * 6) + 1

function myButtonClicked() {
  //input
  const numberGuessed = parseInt(document.getElementById("numberGuessed").value)

  //process
  if (numberGuessed == randomNumber) {
    //output
    document.getElementById("answer").innerHTML = "You guessed correctly!"
  }
  if (numberGuessed != randomNumber) {
    //output
    document.getElementById("answer").innerHTML = "You guessed too high!"
  }
}
