import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { FormBuilder } from '@angular/forms';

import { ApiCallerService } from '../api-caller.service';

export interface IUserData {
  username: string
  password: string
}

@Component({
  selector: 'app-auth-form',
  templateUrl: './auth-form.component.html',
  styleUrls: ['./auth-form.component.sass']
})
export class AuthFormComponent implements OnInit {

  userData:IUserData = {
    username: "",
    password: ""
  }
  ngOnInit(): void {
  }

  signupForm = this.formBuilder.group({
    username: '',
    password: ''
  });

  constructor(private apiCallerService: ApiCallerService, private formBuilder: FormBuilder) {

  }

  onSubmit(): void {
    var userInput = this.signupForm.value
    if(userInput.username && userInput.password) {
      this.signUpUser(userInput.username, userInput.password)
      console.log("Successfully signed up!")
    }
    else {
      console.warn("Something went wrong.")
    }
  }

  signUpUser(username: string, password: string) {
    this.apiCallerService.signUp(username, password)
  }

}
