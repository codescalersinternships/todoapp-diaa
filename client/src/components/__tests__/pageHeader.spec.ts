import PageHeaderVue from '../PageHeader.vue'

describe('Page Header', () => {
  beforeEach(() => cy.mount(PageHeaderVue))

  it('should have input field', () => {
    cy.get('#title').should('exist')
  })

  it('should have Add button', () => {

    cy.get('.add-btn').should('exist')
  })

 
})
