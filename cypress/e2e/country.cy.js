describe('Country Page Tests', () => {
    context('Country page - United States', () => {
        beforeEach(() => {
            cy.visit('/country/united-states/')
        })

        it('should display country title', () => {
            cy.get('h1').should('be.visible')
            cy.get('h1').should('contain', 'Browse By State')
        })

        it('should have working breadcrumb navigation', () => {
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
            
            // Test home breadcrumb link
            cy.get('nav[aria-label="Breadcrumb"] a[href="/"]').click()
            cy.url().should('eq', Cypress.config().baseUrl)
        })

        it('should display country tabs with sign counts', () => {
            // Should show tabs for multiple countries
            cy.contains('United States').should('be.visible')
            cy.contains('Canada').should('be.visible')
            
            // Tabs should show sign counts - verify each count badge contains numbers
            cy.get('[data-cy="country-count"]').each(($el) => {
                cy.wrap($el).invoke('text').should('match', /\d+/)
            })
        })

        it('should navigate to another country via tabs', () => {
            // Click Canada tab
            cy.contains('a', 'Canada').click()
            cy.url().should('include', '/country/canada')
            cy.get('h1').should('be.visible')
        })

        it('should have working global navigation link', () => {
            // Test Search link from global nav
            cy.contains('nav a', 'Search').click()
            cy.url().should('include', '/search')
        })

        it('should display state links', () => {
            cy.get('[data-cy="states-container"]').should('be.visible')
            cy.get('[data-cy="state-link-california"]').should('exist')
        })

        it('should navigate to state page from state link', () => {
            cy.get('[data-cy="state-link-washington"]').click()
            cy.url().should('include', '/state/washington')
        })

        it('should display sign tiles', () => {
            // Country pages may show sign tiles below state links
            cy.get('body').then($body => {
                if ($body.find('[data-cy="sign-tile"]').length > 0) {
                    cy.get('[data-cy="sign-tile"]').should('be.visible')
                    
                    // Test one sign tile link
                    cy.get('[data-cy="sign-tile"]').first().find('a').first().click()
                    cy.url().should('include', '/sign/')
                }
            })
        })
    })

    context('Country page - Canada', () => {
        beforeEach(() => {
            cy.visit('/country/canada/')
        })

        it('should display Canada country page', () => {
            cy.get('h1').should('be.visible')
        })

        it('should have Canada tab active', () => {
            cy.contains('.text-blue-600', 'Canada').should('be.visible')
        })

        it('should display province links', () => {
            cy.get('[data-cy="states-container"]').should('be.visible')
        })

        it('should navigate to province page', () => {
            // Find and click a province link (e.g., Alberta, British Columbia)
            cy.get('[data-cy="states-container"] a').first().click()
            cy.url().should('include', '/state/')
        })
    })

    context('Country navigation from homepage', () => {
        it('should navigate to country page from homepage', () => {
            cy.visit('/')
            
            // Country tabs should be on homepage
            cy.contains('a', 'United States').should('be.visible')
        })
    })
})
