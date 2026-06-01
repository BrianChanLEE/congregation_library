import { render, screen, waitFor } from '@testing-library/react';
import { MovementHistoryScreen } from '../MovementHistoryScreen';
import { vi, describe, it, expect } from 'vitest';

const useMovementHistoryMock = vi.fn();

vi.mock('../../hooks/useMovementHistory', () => ({
  useMovementHistory: () => useMovementHistoryMock(),
}));

describe('MovementHistoryScreen', () => {
  it('renders correctly and handles empty API response', async () => {
    useMovementHistoryMock.mockReturnValue({ data: [], isError: false });

    render(<MovementHistoryScreen />);

    expect(screen.getByText('물품 이동 기록')).toBeInTheDocument();

    await waitFor(() => {
      expect(screen.queryByRole('table')).toBeInTheDocument();
    });
  });

  it('renders transactions correctly', async () => {
    const mockTransactions = [
      { id: 1, time: '2026-05-29 10:00', item: 'Book A', qty: 1, user: 'Admin', memo: 'Test', type: 'IN' },
    ];
    useMovementHistoryMock.mockReturnValue({ data: mockTransactions, isError: false });

    render(<MovementHistoryScreen />);

    await waitFor(() => {
      expect(screen.getByText('Book A')).toBeInTheDocument();
    });
  });
});
