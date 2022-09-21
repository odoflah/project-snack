import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

import { environment } from 'src/environments/environment';
import { Sighting } from './_interfaces/sighting';

const SIGNIN_URL = environment.apiURL + "/auth/signin";
const SIGNUP_URL = environment.apiURL + "/auth/signup";

const SUBMIT_SIGHTING = environment.apiURL + "/snacktrack/submit-sighting"
const GET_SIGHTINGS = environment.apiURL + "/snacktrack/get-sightings"

@Injectable({
  providedIn: 'root'
})
export class ApiCallerService {

  constructor(private http: HttpClient) {
  }

  signIn(username: string, password: string) {
    var httpResponse: string = ""
    this.http.post<any>(SIGNIN_URL, { username: username, password: password })
      .subscribe()
  }

  signUp(username: string, password: string) {
    var httpResponse: string = ""
    this.http.post<any>(SIGNUP_URL, { username: username, password: password }, { observe: 'response' })
      .subscribe(response => {

        // You can access status:
        console.log(response.status)

        // Or any other header:
        console.log(response.headers.get('X-Custom-Header'))
      })
  }

  getSightings(): Observable<Sighting[]> {
    return this.http.get<Sighting[]>(GET_SIGHTINGS)
  }

  submitSighting(sname: string, sighttime: string, sightlocation: string, simage: string, sighter: string): void {
    console.log(sighttime)
    this.http.post<any>(SUBMIT_SIGHTING, { sname: sname, sighttime: sighttime, simage: simage, sightlocation: sightlocation, sighter: sighter }, {}).subscribe(res => {
      console.log(res)
    })
  }
}
