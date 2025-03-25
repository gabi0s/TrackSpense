import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink, RouterOutlet } from '@angular/router';
import { RouterModule } from '@angular/router';
import { AuthService } from './services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  imports: [CommonModule,RouterOutlet,RouterLink],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'app-budget';
  constructor(
    public authService: AuthService,
    private router: Router // Injectez Router
    
  ) {}

  logout() {
    this.authService.logout(); // appelle la méthode logout() du service
    this.router.navigate(['/login']); // redirige vers la page de login
  }
}



