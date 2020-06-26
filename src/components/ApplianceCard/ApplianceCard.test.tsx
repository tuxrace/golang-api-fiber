import React from 'react';
import { render } from '@testing-library/react';
import ApplianceCard from './ApplianceCard';

describe('ApplianceCard', () => {
    test('ApplianceCard renders', () => {   
        const data = { serialNumber: 1111, brand: 'Panasonic', model: 'Tub', status: 'Sold' };
        const {container, getByText} = render(<ApplianceCard data={data}/>);
        expect(container).toBeDefined();
        expect(getByText(/Panasonic/i)).toBeInTheDocument();
    })
})

