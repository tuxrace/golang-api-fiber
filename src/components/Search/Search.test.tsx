import React from 'react';
import { render } from '@testing-library/react';
import Search from './Search';

describe('Search', () => {
    test('Search renders', () => {   
        const props = {
            handleSeach: jest.fn()
        }
        const { container } = render(<Search  {...props}/>);
        expect(container).toBeDefined();
    })
})
