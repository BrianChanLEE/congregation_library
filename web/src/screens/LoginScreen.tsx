import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import { jwtDecode } from 'jwt-decode';
import logo from '/logo.png';
import { apiClient } from '../lib/axios';
import { useAuthStore, type UserRole } from '../store/authStore';

const loginSchema = z.object({
  congId: z.string().min(1, '필수 항목입니다.'),
  email: z.string().email('올바른 이메일 형식이 아닙니다.'),
  password: z.string().min(1, '비밀번호를 입력하세요.'),
});

type LoginFormData = z.infer<typeof loginSchema>;

interface JwtPayload {
  userId: number;
  role: UserRole;
  email: string;
}

export const LoginScreen: React.FC = () => {
  const navigate = useNavigate();
  const { register, handleSubmit, formState: { errors } } = useForm<LoginFormData>({
    resolver: zodResolver(loginSchema),
  });

  const onSubmit = async (data: LoginFormData) => {
    try {
      const response = await apiClient.post('/auth/login', {
        cong_code: data.congId,
        jwhub_email: data.email,
        password: data.password,
      });

      console.log('Login Response:', response.data);

      const { token } = response.data;
      
      if (token) {
        // Decode JWT to extract user info
        const decoded = jwtDecode<JwtPayload>(token);
        console.log('Decoded Token:', decoded);
        
        useAuthStore.getState().setAuth(token, decoded.role || 'user', decoded.userId);
        navigate(decoded.role === 'admin' ? '/admin' : '/');
      } else {
        alert('인증 토큰이 없습니다.');
      }
    } catch (err) {
      console.error('로그인 실패:', err);
      alert('로그인에 실패했습니다.');
    }
  };

  return (
    <main className="w-full min-h-screen flex flex-col md:flex-row bg-background">
      {/* Left Side: Branding */}
      <div className="relative hidden md:flex md:w-1/2 bg-kingdom-blue items-center justify-center overflow-hidden">
        <div className="absolute inset-0 z-0">
          <img alt="Office" className="w-full h-full object-cover opacity-30" src="https://lh3.googleusercontent.com/aida-public/AB6AXuCMmoJuj00W5HzqU9eK0E_REVtyDLqYbLgMLdb3i8oaQxtJ2ZFUAuCbOd4dYpUePIKmc654ur_h16I2pLxZLw8AU9HyxO3-_AhnHlNf9TLH8vms0aBnB7jqAR4KqX3bDXH5y7GqBwTBMe-jkh04rICKQPsckLmWc1KyljqnYzHeu0SWXLdWsBP38yrfy4ygmH5t9mTo5S9Or6r3j8tsJuQk7k6aPe3Q60fdLSoC0CTt5UMe_Up-rOdd578gt8Ws3Y4cifEqLLOnM-w" />
        </div>
        <div className="relative z-10 text-center px-12">
          {/* Logo size adjusted to match stitch_ref */}
          <img src={logo} alt="Logo" className="w-[104px] h-[104px] rounded-2xl mb-6 shadow-xl border border-white/20 mx-auto" />
          <h1 className="font-headline-lg text-white text-5xl tracking-tight mb-4">Kingdom Ledger</h1>
          <p className="font-body-lg text-white/80 max-w-md mx-auto">The complete digital ecosystem for managing congregation literature and resources with precision and ease.</p>
        </div>
      </div>

      {/* Right Side: Login Form */}
      <div className="w-full md:w-1/2 flex items-center justify-center bg-white p-6 md:p-12 overflow-y-auto">
        <div className="w-full max-w-[440px] py-12">
          <div className="text-center mb-8 md:hidden">
            <img src={logo} alt="Logo" className="w-16 h-16 rounded-xl mb-4 shadow-sm mx-auto" />
            <h1 className="font-headline-lg-mobile text-kingdom-blue">Kingdom Ledger</h1>
          </div>
          
          <div className="mb-6"> {/* Reduced margin bottom to match screen.png */}
            <h2 className="font-headline-md text-on-surface mb-2">Welcome back</h2>
            <p className="font-body-md text-on-surface-variant">Please enter your details to sign in.</p>
          </div>

          <form onSubmit={handleSubmit(onSubmit)} className="space-y-4"> {/* Reduced gap-5 to space-y-4 for tighter spacing */}
            {[
              { id: 'congId', label: 'CONGREGATION ID', type: 'number', icon: 'pin', placeholder: 'e.g. 12345' },
              { id: 'email', label: 'JWPUB EMAIL', type: 'email', icon: 'mail', placeholder: 'username@jwpub.org' },
              { id: 'password', label: 'PASSWORD', type: 'password', icon: 'lock', placeholder: '••••••••' },
            ].map(field => (
              <div key={field.id}>
                <label className="block font-label-caps text-gray-800 mb-2">{field.label}</label> {/* Changed text color to gray-800 */}
                <div className="relative">
                  <span className="absolute inset-y-0 left-0 pl-3 flex items-center text-outline">
                    <span className="material-symbols-outlined text-[20px]">{field.icon}</span>
                  </span>
                  <input
                    {...register(field.id as keyof LoginFormData)}
                    type={field.type}
                    className="block w-full pl-10 pr-4 py-3 bg-surface border border-outline-variant rounded-lg font-body-md text-on-surface focus:ring-2 focus:ring-kingdom-blue outline-none transition-all"
                    placeholder={field.placeholder}
                  />
                </div>
                {errors[field.id as keyof LoginFormData] && <p className="text-error text-xs mt-1">{errors[field.id as keyof LoginFormData]?.message}</p>}
              </div>
            ))}
            
            <button type="submit" className="w-full bg-kingdom-blue text-white font-headline-sm py-4 rounded-lg shadow-md hover:bg-primary transition-all active:scale-[0.98] mt-4 flex items-center justify-center gap-2">
              Sign In <span className="material-symbols-outlined text-[20px]">arrow_forward</span>
            </button>
          </form>

          <div className="text-center my-8 font-label-caps text-on-surface-variant">DON'T HAVE AN ACCOUNT?</div>
          
          <div className="text-center mb-8">
            <Link to="/register" className="font-headline-sm text-kingdom-blue hover:underline flex items-center justify-center gap-1">
              Create Account <span className="material-symbols-outlined text-[18px]">chevron_right</span>
            </Link>
          </div>
        </div>
      </div>
    </main>
  );
};

export default LoginScreen;
