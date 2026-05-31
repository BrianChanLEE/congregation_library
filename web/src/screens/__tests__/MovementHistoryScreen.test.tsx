import { render, screen, waitFor } from '@testing-library/react';
import { MovementHistoryScreen } from '../MovementHistoryScreen';
import { apiClient } from '../../lib/axios';
import { vi, describe, it, expect } from 'vitest';
import { AxiosResponse } from 'axios';

vi.mock('../../lib/axios', () => ({
  apiClient: {
    get: vi.fn(),
  },
}));

describe('MovementHistoryScreen', () => {
  it('renders correctly and handles empty API response', async () => {
    (apiClient.get as vi.Mock).mockResolvedValue({ data: null } as AxiosResponse);

    render(<MovementHistoryScreen />);
    
    // Check if the title is rendered
    expect(screen.getByText('물품 이동 기록')).toBeInTheDocument();

    // Ensure it doesn't crash
    await waitFor(() => {
        expect(screen.queryByRole('table')).toBeInTheDocument();
    });
  });

  it('renders transactions correctly', async () => {
      const mockTransactions = [
          { id: 1, time: '2026-05-29 10:00', item: 'Book A', qty: 1, user: 'Admin', memo: 'Test', type: 'IN' }
      ];
      (apiClient.get as vi.Mock).mockResolvedValue({ data: mockTransactions } as AxiosResponse);

      render(<MovementHistoryScreen />);

      await waitFor(() => {
          expect(screen.getByText('Book A')).toBeInTheDocument();
      });
  });
});
