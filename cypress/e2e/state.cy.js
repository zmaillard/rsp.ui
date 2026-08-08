describe('State Page Tests', () => {
    context('State page with full content - Washington', () => {
        beforeEach(() => {
            cy.visit('/state/washington/')
        })

        it('should display state title and metadata', () => {
            cy.get('h1').should('be.visible')
            cy.get('h1').should('contain', 'Washington')
        })

        it('should have working breadcrumb navigation', () => {
            cy.get('[data-cy="breadcrumb"]').should('be.visible')
            cy.get('[data-cy="breadcrumb-home"]').should('be.visible')
            
            // Test country breadcrumb link
            cy.get('[data-cy="breadcrumb-country"]').should('be.visible')
            cy.get('[data-cy="breadcrumb-country"]').click()
            cy.url().should('include', '/country/')
            cy.go('back')
        })

        it('should have working global navigation link', () => {
            // Test Recent link from global nav
            cy.contains('a', 'Recent').click()
            cy.url().should('include', '/recent')
        })

        it('should display sign tiles and navigate to sign page', () => {
            cy.get('[data-cy="sign-tile"]').should('be.visible')
            cy.get('[data-cy="sign-tile"]').should('have.length.at.least', 1)
            
            // Test one sign tile link
            cy.get('[data-cy="sign-tile"]').first().find('a').first().click()
            cy.url().should('include', '/sign/')
        })

        it('should display and navigate to county', () => {
            // County links appear in sign metadata
            cy.get('[data-cy="county-link"]').should('exist')
            cy.get('[data-cy="county-link"]').first().click()
            cy.url().should('include', '/county/')
        })

        it('should display and navigate to highway', () => {
            // Highway shield links appear in sign tiles
            cy.get('[data-cy="highway-link"]').should('exist')
            cy.get('[data-cy="highway-link"]').first().click()
            cy.url().should('include', '/highway/')
        })

        it('should display and navigate to place', () => {
            // Place links appear in sign metadata
            cy.get('[data-cy="place-link"]').should('exist')
            cy.get('[data-cy="place-link"]').first().click()
            cy.url().should('include', '/place/')
        })

        it('should have pagination', () => {
            cy.get('[data-cy="pagination"]').should('be.visible')
            cy.get('[data-cy="pagination-next"]').should('be.visible')
            
            // Test pagination navigation
            cy.get('[data-cy="pagination-next"]').click()
            cy.url().should('include', '/page/')
        })
    })

    context('State page - California', () => {
        beforeEach(() => {
            cy.visit('/state/california/')
        })

        it('should display California state title', () => {
            cy.get('h1').should('contain', 'California')
        })

        it('should have sign content', () => {
            cy.get('[data-cy="sign-tile"]').should('have.length.at.least', 1)
        })

        it('should have working breadcrumb to country', () => {
            cy.get('[data-cy="breadcrumb-country"]').should('be.visible')
            cy.get('[data-cy="breadcrumb-country"]').should('contain', 'United States')
        })
    })

    context('State navigation from homepage', () => {
        it('should navigate to state page from homepage state link', () => {
            cy.visit('/')
            cy.get('[data-cy="state-link-washington"]').click()
            cy.url().should('include', '/state/washington')
            cy.get('h1').should('contain', 'Washington')
        })
    })
})
