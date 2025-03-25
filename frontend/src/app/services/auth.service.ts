import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject } from 'rxjs';
import { tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private isConnectedSubject = new BehaviorSubject<boolean>(false);
  isConnected$ = this.isConnectedSubject.asObservable();
  private apiUrl = 'http://localhost:8080/auth';

  constructor(private http: HttpClient) {}


  isAuthenticated(): boolean {
    return !!localStorage.getItem('user_id');
  }

  login(nom: string, password: string) {
    return this.http.post<{ user_id: string }>(this.apiUrl, { Nom: nom, Password: password }).pipe(
      tap((response) => {
        localStorage.setItem('user_id', response.user_id);
        this.isConnectedSubject.next(true);
      })
    );
  }

  logout() {
    localStorage.removeItem('user_id');
    this.isConnectedSubject.next(false);
  }

  getUserId(): string | null {
    return localStorage.getItem('user_id'); // Récupérer user_id du localStorage
  }
}