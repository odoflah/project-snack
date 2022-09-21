import { Component, OnInit } from '@angular/core';

import { ApiCallerService } from '../api-caller.service';
import { Sighting } from '../_interfaces/sighting';

// export interface Snack {
//   name: string;
//   location: string;
//   picture: string;
// }

const ELEMENT_DATA: Sighting[] = [
  { sname: 'Kit Kat', simage: "", sighttime: "2022-09-21 11:18:46.473756", sightlocation: "Velasco mk4-6" },
];


@Component({
  selector: 'app-snack-list',
  templateUrl: './snack-list.component.html',
  styleUrls: ['./snack-list.component.sass', '../app.component.sass']
})
export class SnackListComponent implements OnInit {

  displayedColumns: string[] = ['sname', 'sightlocation'];
  dataSource: Sighting[] = ELEMENT_DATA;


  constructor(private apiCallerService: ApiCallerService) {
  }

  ngOnInit(): void {
    this.apiCallerService.getSightings().subscribe((data: Sighting[]) => {
      console.log(data)
      this.dataSource = data
      console.log(this.dataSource)
    })
  }
}
