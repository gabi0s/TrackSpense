import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { ExpensesService } from '../services/expenses.service';
import { BudgetService } from '../services/budget.service';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-expenses',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './expenses.component.html',
  styleUrls: ['./expenses.component.css']
})
export class ExpensesComponent implements OnInit {
  categories: any[] = [];
  selectedCategory: any = null;
  expenseAmount = 0;
  expenses: any[] = [];
  userId: string | null = null;
  isLoading = false;
  errorMessage = '';

  constructor(
    private expensesService: ExpensesService,
    private budgetService: BudgetService,
    private authService: AuthService
  ) {}

  ngOnInit(): void {
    this.userId = this.authService.getUserId();
    if (!this.userId) {
      this.errorMessage = 'Utilisateur non connecté';
      return;
    }
    this.loadData();
  }

  loadData(): void {
    this.loadCategories();
    this.loadExpenses();
  }

  loadCategories(): void {
    this.isLoading = true;
    this.errorMessage = '';
    
    if (!this.userId) {
      this.errorMessage = 'ID utilisateur non disponible';
      this.isLoading = false;
      return;
    }

    this.budgetService.getBudgets(this.userId).subscribe({
      next: (data) => {
        console.log('Budgets reçus:', data); // Debug: vérifiez la structure des données
        this.categories = data.map(budget => ({
          budget_id: budget.budgets_id, // Selon ce que renvoie l'API
          category: budget.category,
          budgets_limit: budget.budgets_limit
        }));
        this.isLoading = false;
      },
      error: (error) => {
        console.error('Error loading categories:', error);
        this.errorMessage = 'Erreur lors du chargement des catégories';
        this.isLoading = false;
      }
    });
  }

  loadExpenses(): void {
    if (!this.userId) return;

    this.isLoading = true;
    this.expensesService.getExpenses(this.userId).subscribe({
      next: (data) => {
        console.log('Expenses:', data);
        this.expenses = data;
        this.isLoading = false;
      },
      error: (error) => {
        console.error('Error loading expenses:', error);
        this.errorMessage = 'Erreur lors du chargement des dépenses';
        this.isLoading = false;
      }
    });
  }

  addExpense(): void {
    if (!this.userId) {
      this.errorMessage = 'Utilisateur non connecté';
      return;
    }

    if (!this.selectedCategory || !this.selectedCategory.budget_id || this.expenseAmount <= 0) {
      this.errorMessage = 'Veuillez sélectionner une catégorie valide et entrer un montant valide';
      return;
    }

    this.isLoading = true;
    this.errorMessage = '';

    console.log('Envoi de la dépense avec:', { // Debug
      userId: this.userId,
      budgetId: this.selectedCategory.budget_id,
      category: this.selectedCategory.category,
      amount: this.expenseAmount
    });

    this.expensesService.createExpense(
      this.userId,
      this.selectedCategory.budget_id,
      this.selectedCategory.category,
      this.expenseAmount
    ).subscribe({
      next: (response) => {
        console.log('Réponse du serveur:', response);
        this.loadExpenses();
        this.expenseAmount = 0;
        this.selectedCategory = null;
        this.isLoading = false;
      },
      error: (error) => {
        console.error('Error adding expense:', error);
        this.errorMessage = 'Erreur lors de l\'ajout de la dépense: ' + 
          (error.error?.message || error.message);
        this.isLoading = false;
      }
    });
  }
}