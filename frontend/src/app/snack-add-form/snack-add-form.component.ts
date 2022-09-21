import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { Sighting } from '../_interfaces/sighting';

@Component({
  selector: 'app-snack-add-form',
  templateUrl: './snack-add-form.component.html',
  styleUrls: ['./snack-add-form.component.sass']
})
export class SnackAddFormComponent implements OnInit {

  constructor(private formBuilder: FormBuilder) { }

  ngOnInit(): void {
  }

  snackSubmission = this.formBuilder.group({
    name: "",
    location: "",
    sighter: "",
  });

  newSnackSighting(): void {
    let formVal = this.snackSubmission.value
    console.log(formVal)
    // let newSighting: Sighting = {
    //   snack: formVal.name ?? undefined,
    //   location: this.snackSubmission.value.location,
    //   picture: null,
    //   sighter: this.snackSubmission.value.sighter,
    //   timestamp: Date.now()
    // }
    console.log("New snack")

  }

}
