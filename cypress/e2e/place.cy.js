describe('Place Page Tests', () => {
    context('Place page - Seattle, Washington', () => {
        beforeEach(() => {
            cy.visit('/place/washington_seattle/')
        })

        it('should display place title with city and state', () => {
            cy.get('h1').should('be.visible')
            cy.get('h1').should('contain', 'Seattle')
            cy.get('h1').should('contain', 'Washington')
        })

        it('should have working breadcrumb navigation', () => {
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
            
            // Test state breadcrumb link
            cy.contains('a', 'Washington').click()
            cy.url().should('include', '/state/washington')
            cy.go('back')
            
            // Test country breadcrumb link
            cy.contains('a', 'United States').click()
            cy.url().should('include', '/country/united-states')
        })

        it('should have working global navigation link', () => {
            cy.visit('/place/washington_seattle/')
            
            // Test Home link from global nav
            cy.contains('nav a', 'Home').click()
            cy.url().should('eq', Cypress.config().baseUrl)
        })

        it('should display sign tiles and navigate to sign page', () => {
            cy.get('[data-cy="sign-tile"]').should('be.visible')
            cy.get('[data-cy="sign-tile"]').should('have.length.at.least', 1)
            
            // Test one sign tile link
            cy.get('[data-cy="sign-tile"]').first().find('a').first().click()
            cy.url().should('include', '/sign/')
        })
    })

    context('Place page - Foley, Alabama', () => {
        beforeEach(() => {
            cy.visit('/place/alabama_foley/')
        })

        it('should display place title correctly', () => {
            cy.get('h1').should('contain', 'Foley')
            cy.get('h1').should('contain', 'Alabama')
        })

        it('should have sign content', () => {
            cy.get('[data-cy="sign-tile"]').should('have.length.at.least', 1)
        })

        it('should have correct breadcrumb hierarchy', () => {
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
            cy.contains('a', 'Alabama').should('be.visible')
            cy.contains('a', 'United States').should('be.visible')
        })
    })

    context('Navigation to place from state page', () => {
        it('should navigate to place page from state page place link', () => {
            cy.visit('/state/washington/')
            
            // Click a place link from sign metadata
            cy.get('[data-cy="place-link"]').first().then($link => {
                const placeName = $link.text()
                let places = placeName.trim().split('\n');
                cy.wrap($link).click()
                
                cy.url().should('include', '/place/')
                cy.get('h1').should('contain', places[0])
            })
        })
    })
})
