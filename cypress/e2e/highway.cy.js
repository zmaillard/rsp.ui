describe('Highway Page Tests', () => {
    context('Highway without external link (id44)', () => {
        beforeEach(() => {
            cy.visit('/highway/id44')
        })

        it('should display the highway title and shield image', () => {
            cy.get('[data-cy="highway-shield"]').should('be.visible')
            cy.get('[data-cy="highway-title"]').should('be.visible')
            cy.get('[data-cy="highway-title"]').parent('a').should('not.exist')
        })

        it('should display the correct breadcrumb navigation', () => {
            // Check breadcrumb structure
            cy.get('[data-cy="breadcrumb"]').should('be.visible')
            cy.get('[data-cy="breadcrumb-home"]').should('be.visible')
            cy.get('[data-cy="breadcrumb-country"]').should('be.visible')
            cy.get('[data-cy="breadcrumb-highway-type"]').should('be.visible')
            cy.get('[data-cy="breadcrumb-current"]').should('be.visible')

            // Test navigation
            cy.get('[data-cy="breadcrumb-home"]').click()
            cy.url().should('eq', Cypress.config().baseUrl )
            cy.go('back')

            cy.get('[data-cy="breadcrumb-country"]').click()
            cy.url().should('include', '/country/')
            cy.go('back')

            cy.get('[data-cy="breadcrumb-highway-type"]').click()
            cy.url().should('include', '/highwaytype/')
        })

        it('should display sign tiles or features', () => {
            // Check if features exist
            cy.get('body').then($body => {
                if ($body.find('[data-cy="feature-summary"]').length > 0) {
                    // Test feature summaries
                    cy.get('[data-cy="feature-summary"]').should('be.visible')
                    cy.get('[data-cy="feature-summary"]').first().children('a').click()
                    cy.url().should('include', '/feature/')
                } else {
                    // Test sign tiles
                    cy.get('[data-cy="sign-tile"]').should('be.visible')
                    cy.get('[data-cy="sign-tile"]').first().click()
                    cy.url().should('include', '/sign/')
                }
            })
        })
    })

    context('Highway with external link (az195)', () => {
        beforeEach(() => {
            cy.visit('/highway/az195')
        })

        it('should display the highway title with external link', () => {
            cy.get('[data-cy="highway-shield"]').should('be.visible')
            cy.get('[data-cy="highway-title"]').should('be.visible')
            cy.get('[data-cy="highway-title"]').parent('a').should('have.attr', 'href')
            cy.get('[data-cy="highway-title"]').parent().children('svg').should('be.visible').should('have.class', 'dark:text-white')
        })

        it('should open external link in new tab', () => {
            cy.get('[data-cy="highway-title"]').parent('a').should('have.attr', 'target', '_blank')
            cy.get('[data-cy="highway-title"]').parent('a').should('have.attr', 'rel', 'noopener noreferrer')
        })
    })

    context('Multi-state highways', () => {
        beforeEach(() => {
            // Visit a highway that spans multiple states, like i90
            cy.visit('/highway/i90')
        })

        it('should display state headings when highway spans multiple states', () => {
            cy.get('[data-cy="state-links"]').should('be.visible')
            cy.get('[data-cy="state-heading"]').should('have.length.at.least', 2)
        })

        it('should navigate to state sections when clicking state links', () => {
            cy.get('[data-cy="state-heading"]').first().then($stateLink => {
                const stateId = $stateLink.attr('href').replace('#', '')

                cy.get('[data-cy="state-heading"]').first().click()

                // URL should have the hash
                cy.hash().should('eq', `#${stateId}`)

                // Page should scroll to that section
                cy.get(`#${stateId}`).should('be.visible')
            })
        })

        it('should group features by state', () => {
            cy.get('[data-cy="state-heading"]').each(($heading) => {
                // Get the state section
                const stateId = $heading.attr('href').replace('#', '')

                // Each state should have features
                cy.get(`#${stateId}`).parent().find('[data-cy="feature-summary"]')
                    .should('be.visible')
            })
        })
    })
})