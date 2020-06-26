import React from 'react';
import { render } from '@testing-library/react';
import AppliancesList from './AppliancesList';

describe('AppliancesList', () => {
    test('AppliancesList renders', () => {   
        const { container } = render(<AppliancesList />);
        expect(container).toBeDefined();
    })
    test('No Records Found renders', () => {   
        const { getByText } = render(<AppliancesList />);
        expect(getByText(/No Records Found/i)).toBeInTheDocument();
    })
})
