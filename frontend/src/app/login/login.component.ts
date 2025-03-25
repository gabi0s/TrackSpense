import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { AuthService } from '../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [FormsModule, CommonModule],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  email: string = '';
  password: string = '';
  isConnected: boolean = false; // connection state

  constructor(private authService: AuthService, private router: Router) {}

  // forms submision
  onSubmit() {
    this.authService.login(this.email, this.password).subscribe({
      next: () => {
        const userId = this.authService.getUserId();
        console.log('User ID:', userId);
        setTimeout(() => {
          this.router.navigate(['/budget']);
        }, 1000);
      },
      error: () => {
        alert('Identifiants incorrects');
      }
    });
  }
}