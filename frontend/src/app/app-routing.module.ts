import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { SnackListComponent } from './snack-list/snack-list.component';
import { SnackAddFormComponent } from './snack-add-form/snack-add-form.component';

const routes: Routes = [
  { path: '', component: SnackListComponent },
  { path: 'add-snack', component: SnackAddFormComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
