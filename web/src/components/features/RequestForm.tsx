import React from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import { apiClient } from '../../lib/axios';

// DESIGN.md 기반 수량 선택 폼
const requestSchema = z.object({
  quantity: z.number().min(1, '최소 1개 이상').max(99, '최대 99개'),
});

type RequestFormData = z.infer<typeof requestSchema>;

interface RequestFormProps {
  itemId: number;
  toCongId: number;
  onSuccess: () => void;
}

export const RequestForm: React.FC<RequestFormProps> = ({ itemId, toCongId, onSuccess }) => {
  const {
    register,
    handleSubmit,
    setValue,
    watch,
    formState: { errors, isSubmitting },
  } = useForm<RequestFormData>({
    resolver: zodResolver(requestSchema),
    defaultValues: { quantity: 1 },
  });

  const quantity = watch('quantity');

  const onSubmit = async (data: RequestFormData) => {
    try {
      await apiClient.post('/transactions', {
        to_cong_id: toCongId,
        item_id: itemId,
        quantity: -data.quantity,
        type: 'OUT',
      });
      onSuccess();
    } catch (err) {
      console.error('신청 실패:', err);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-md">
      <div className="flex flex-col gap-sm">
        <label className="font-label-md text-on-surface-variant uppercase tracking-wider">신청 수량</label>
        <div className="flex items-center border border-outline-variant rounded-default overflow-hidden h-12">
          <button 
            type="button"
            className="w-12 h-full bg-surface-container-low hover:bg-surface-container transition-colors flex items-center justify-center text-primary"
            onClick={() => setValue('quantity', Math.max(1, quantity - 1))}
          >
            <span className="material-symbols-outlined">remove</span>
          </button>
          <input 
            className="flex-1 text-center border-none h-full focus:ring-0 font-body-md text-primary"
            type="number" 
            {...register('quantity', { valueAsNumber: true })}
          />
          <button 
            type="button"
            className="w-12 h-full bg-surface-container-low hover:bg-surface-container transition-colors flex items-center justify-center text-primary"
            onClick={() => setValue('quantity', quantity + 1)}
          >
            <span className="material-symbols-outlined">add</span>
          </button>
        </div>
        {errors.quantity && <p className="font-body-sm text-error">{errors.quantity.message}</p>}
      </div>
      <button 
        type="submit"
        disabled={isSubmitting}
        className="w-full py-md bg-primary text-on-primary rounded-default font-label-md hover:brightness-110 active:scale-95 transition-all mt-sm"
      >
        {isSubmitting ? '신청 중...' : '신청하기'}
      </button>
    </form>
  );
};
