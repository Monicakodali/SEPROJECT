import React from 'react'
import { mount } from '@cypress/react'
import App from '../App'

beforeEach(() => {
  mount(<App />)
  cy.routerNavigate('/signup')
})

it('App loads', () => {
  cy.get('main').should('be.visible')
})