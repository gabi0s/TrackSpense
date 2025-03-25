import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BudgetService {
  private apiUrl = 'http://localhost:8082'; // Base URL de l'API

  constructor(private http: HttpClient) {}

  // 1️ Créer un budget
  createBudget(userId: string, category: string, budgetLimit: number): Observable<any> {
    const body = { user_id: userId, category: category, budgets_limit: budgetLimit };
    return this.http.post<any>(`${this.apiUrl}/create-budget`, body);
  }

  // 2️ Récupérer tous les budgets d'un utilisateur
  getBudgets(userId: string): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/get-budget/${userId}`);
  }


  // 4️ Mettre à jour un budget
  updateBudget(userId: string, budgetId: string, amount: number): Observable<any> {
    const body = { amount: amount };
    return this.http.put<any>(`${this.apiUrl}/update-budget/${userId}/${budgetId}`, body);
  }

  // 5️ Supprimer un budget
  deleteBudget(userId: string, budgetId: string): Observable<any> {
    return this.http.delete<any>(`${this.apiUrl}/delete-budget/${userId}/${budgetId}`);
  }
}
