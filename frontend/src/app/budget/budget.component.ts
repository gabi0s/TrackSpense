import { Component, OnInit } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { SoldeService } from '../services/solde.service';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { BudgetService } from '../services/budget.service';
import { ExpensesService } from '../services/expenses.service';

@Component({
  selector: 'app-budget',
  standalone: true,
  imports: [FormsModule, CommonModule],
  templateUrl: './budget.component.html',
  styleUrls: ['./budget.component.css']
})
export class BudgetComponent implements OnInit {
  solde: number | null = null;
  userId: string | null = null;
  customAmount: number = 0;
  categories: string[] = [];
  budgets: any[] = [];
  newCategory: string = '';
  budgetLimit: number = 0;
  expenses: any[] = [];

  constructor(
    private authService: AuthService,
    private soldeService: SoldeService,
    private budgetService: BudgetService,
    private expensesService: ExpensesService
  ) {}

  ngOnInit() {
    
    this.userId = this.authService.getUserId();
    console.log('User ID récupéré:', this.userId);

    if (this.userId) {
      this.loadUserData();
    }
  }

  loadUserData() {
    this.soldeService.getSolde(this.userId!).subscribe({
      next: (response) => {
        this.solde = response.solde;
      },
      error: () => {
        console.error('Erreur lors de la récupération du solde');
      }
    });

    this.getUserBudgets();
    this.loadExpenses();
  }

  updateSolde(amount: number) {
    if (this.userId) {
      this.soldeService.updateSolde(this.userId, amount).subscribe({
        next: (response) => {
          this.solde = response.solde;
          this.customAmount = 0;
        },
        error: () => {
          console.error('Erreur lors de la mise à jour du solde');
        }
      });
    }
  }

  createNewBudget() {
    if (this.userId && this.newCategory && this.budgetLimit > 0) {
      this.budgetService.createBudget(this.userId, this.newCategory, this.budgetLimit).subscribe({
        next: (response) => {
          console.log('Budget créé avec succès:', response);
          this.newCategory = '';
          this.budgetLimit = 0;
          this.getUserBudgets();
        },
        error: () => {
          console.error('Erreur lors de la création du budget');
        }
      });
    }
  }

  getUserBudgets() {
    if (this.userId) {
      this.budgetService.getBudgets(this.userId).subscribe({
        next: (response) => {
          console.log('Budgets récupérés:', response);
          this.budgets = response;
          this.categories = Array.from(new Set(response.map(budget => budget.category)));
          this.updateSpentAmounts();
        },
        error: () => {
          console.error('Erreur lors de la récupération des budgets');
        }
      });
    }
  }

  loadExpenses() {
    if (this.userId) {
      this.expensesService.getExpenses(this.userId).subscribe({
        next: (expenses) => {
          this.expenses = expenses;
          this.updateSpentAmounts();
        },
        error: (error) => {
          console.error('Erreur lors du chargement des dépenses:', error);
        }
      });
    }
  }

  updateSpentAmounts() {
    if (this.budgets.length > 0 && this.expenses.length > 0) {
      this.budgets.forEach(budget => {
        budget.current_amount = this.calculateSpentForCategory(budget.category);
      });
    }
  }

  calculateSpentForCategory(category: string): number {
    const categoryExpenses = this.expenses.filter(
      expense => expense.budget_categorie === category
    );
    return categoryExpenses.reduce(
      (total, expense) => total + expense.price, 0
    );
  }

  deleteBudget(budgetId: string) {
    if (this.userId) {
      this.budgetService.deleteBudget(this.userId, budgetId).subscribe({
        next: () => {
          this.getUserBudgets();
        },
        error: (error) => {
          console.error('Erreur lors de la suppression du budget:', error);
        }
      });
    }
  }
}