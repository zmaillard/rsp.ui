describe('Edited Sign Widget Tests', () => {
    context('Edited sign (1654110315)', () => {
        beforeEach(() => {
            cy.visit('/sign/1654110315')
        })

        it('should display the edited text widget', () => {
            cy.get('[data-cy="sign-edited-text"]').should('exist').and('be.visible')
        })

        it('should contain the AI enhancement message', () => {
            cy.get('[data-cy="sign-edited-text"]')
                .should('contain.text', 'This sign has been enhanced by an AI model')
        })

        it('should initially show "View Original" toggle', () => {
            cy.get('#sign-edit-link-toggle').should('have.text', 'View Original')
        })

        it('should toggle to "View Edited" when clicked', () => {
            cy.get('#sign-edit-link').click()
            cy.get('#sign-edit-link-toggle').should('have.text', 'View Edited')
        })

        it('should toggle back to "View Original" when clicked again', () => {
            cy.get('#sign-edit-link').click()
            cy.get('#sign-edit-link-toggle').should('have.text', 'View Edited')
            cy.get('#sign-edit-link').click()
            cy.get('#sign-edit-link-toggle').should('have.text', 'View Original')
        })
    })

    context('Non-edited sign (3406055620)', () => {
        beforeEach(() => {
            cy.visit('/sign/3406055620')
        })

        it('should not display the edited text widget', () => {
            cy.get('[data-cy="sign-edited-text"]').should('not.exist')
        })
    })
})
