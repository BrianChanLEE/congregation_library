// Note 1: Define core types for ledger system.
export interface Item {
  id: number;
  name: string;
  code: string;
  category: string;
  imageUrl: string;
  stock: number;
}
export interface Transaction {
  id: number;
  toCongId: number;
  itemId: number;
  quantity: number;
  type: 'IN' | 'OUT' | 'CANCEL';
  createdAt: string;
}

export interface AdminStats {
  totalItems: number;
  lowStockCount: number;
  pendingRequests: number;
}

export interface Alert {
  message: string;
  severity: 'INFO' | 'WARNING' | 'ERROR';
  time: string;
}

export interface UserProfile {
  id: number;
  name: string;
  position: string;
  phone: string;
  email: string;
  jwhubEmail: string;
}

export interface Inventory {
  congId: number;
  itemId: number;
  quantity: number;
}
