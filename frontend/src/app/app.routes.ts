import { NgModule } from '@angular/core';
import { Routes } from '@angular/router';
import { ExpensesComponent } from './expenses/expenses.component';
import { LoginComponent } from './login/login.component';
import { BudgetComponent } from './budget/budget.component';
import { RouterModule } from '@angular/router';

export const routes: Routes = [
    {path: 'login', component: LoginComponent},
    {path: 'expenses', component: ExpensesComponent},
    {path: 'budget', component: BudgetComponent},
    //{ path: '**', component: NotFoundComponent }

];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
  })
  export class AppRoutingModule { }