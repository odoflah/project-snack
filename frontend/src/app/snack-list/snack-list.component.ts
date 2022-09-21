import { Component, OnInit } from '@angular/core';

export interface Snack {
  name: string;
  location: string;
  picture: string;
}

const ELEMENT_DATA: Snack[] = [
  { name: 'Kit Kat', location: "Docs MK8-3", picture: 'H' },
  { name: 'Orange Chocolate', location: "Docs MK8-3", picture: 'He' },
  { name: 'Smoked Almonds', location: "Docs MK8-3", picture: 'Li' },
  { name: 'Spinach', location: "Docs MK8-3", picture: 'Be' },
  { name: 'Brocoli', location: "Docs MK8-3", picture: 'B' },
  { name: 'Roast peppers', location: "Docs MK8-3", picture: 'C' },
  { name: 'Water', location: "Docs MK8-3", picture: 'N' },
];


@Component({
  selector: 'app-snack-list',
  templateUrl: './snack-list.component.html',
  styleUrls: ['./snack-list.component.sass', '../app.component.sass']
})
export class SnackListComponent implements OnInit {

  displayedColumns: string[] = ['name', 'location', 'picture'];
  dataSource = ELEMENT_DATA;

  constructor() { }

  ngOnInit(): void {
  }



}
