import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthService } from '../auth.service';


@Injectable()
export class TokeninterceptorInterceptor implements HttpInterceptor {

  constructor(public auth: AuthService, ) {}

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    if (this.auth.tokenInUse()) {
      request = request.clone({
        setHeaders: {
          Authorization: `Bearer ${this.auth.getAuthToken()}`
        }
      });
    }
    return next.handle(request);
  }
}
