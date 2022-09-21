import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { Sighting } from '../_interfaces/sighting';
import { ApiCallerService } from '../api-caller.service';

@Component({
  selector: 'app-snack-add-form',
  templateUrl: './snack-add-form.component.html',
  styleUrls: ['./snack-add-form.component.sass', '../app.component.sass']
})
export class SnackAddFormComponent implements OnInit {

  constructor(private formBuilder: FormBuilder, private apiCallerService: ApiCallerService) { }

  ngOnInit(): void {
  }

  snackSubmission = this.formBuilder.group({
    name: "",
    location: "",
    sighter: "",
  });

  newSnackSighting(): void {
    let formVal = this.snackSubmission.value

    this.apiCallerService.submitSighting(formVal.name!, Date.now().toString(), formVal.location!, "", formVal.sighter!)
    console.log(formVal.sighter)
    // let newSighting: Sighting;

    // newSighting.sname = formVal.name

    // sname  formVal.name,
    // simage: "",
    //   sighttime: Date.now().toString(),
    //     sightlocation: formVal.location,
    //       sighter: formVal.sighter

    console.log("New snack")

  }
}

// curl -X POST -v -H "Content-Type: application/json" -d '{"sname": "Kit kat", "sighttime": "2022-09-21 12:18:46.473756", "sightlocation": "ssdfdfgdff", "simage": "", "sighter": "Alex"}' http://localhost:8000/snacktrack/submit-sighting
