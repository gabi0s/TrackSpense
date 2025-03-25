import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class SoldeService {
  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  getSolde(userId: string) {
    return this.http.get<{ message: string; solde: number }>(`${this.apiUrl}/get-solde/${userId}`);
  }

  updateSolde(userId: string, amount: number) {
    return this.http.post<{ message: string; solde: number }>(`${this.apiUrl}/update-solde/`, { user_id: userId, amount });
  }
}