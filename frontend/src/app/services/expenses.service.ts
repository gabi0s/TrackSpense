import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { BudgetService } from '../services/budget.service';
import { SoldeService } from './solde.service';

@Injectable({
  providedIn: 'root'
})
export class ExpensesService {
  private apiUrl = 'http://localhost:8081'; // Base URL de l'API des dépenses

  constructor(
    private http: HttpClient,
    private budgetService: BudgetService,
    private soldeService: SoldeService
  ) {}



  // Créer une nouvelle dépense
  createExpense(userId: string, budgetId: string, category: string, price: number): Observable<any> {
    const body = { 
      user_id: userId, 
      budget_id: budgetId, 
      budget_categorie: category, 
      price: price 
    };
    return new Observable(observer => {
      this.http.post<any>(`${this.apiUrl}/create-expense`, body).subscribe({
        next: (response) => {
          this.soldeService.updateSolde(userId, -price).subscribe({
            next: () => {
              observer.next(response);
              observer.complete();
            },
            error: (error) => {
              observer.error(error);
            }
          });
        },
        error: (error) => {
          observer.error(error);
        }
      });
    });
  }
  

  // Récupérer toutes les dépenses d'un utilisateur
  getExpenses(userId: string): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/get-expense/${userId}`);
  }

  // Récupérer les budgets/catégories disponibles pour un utilisateur
  getBudgetCategories(userId: string): Observable<any[]> {
    return this.budgetService.getBudgets(userId);
  }
}