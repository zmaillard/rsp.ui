describe('Highway Type Page Tests', () => {
    context('Highway type page - Interstate Highway', () => {
        beforeEach(() => {
            cy.visit('/highwaytype/interstate-highway/')
        })

        it('should display highway type title', () => {
            cy.get('[data-cy="highway-type-title"]').should('be.visible')
            cy.get('[data-cy="highway-type-title"]').should('contain', 'Interstate Highway')
        })

        it('should have working breadcrumb navigation', () => {
            cy.get('[data-cy="breadcrumb"]').should('be.visible')
            cy.get('[data-cy="breadcrumb-home"]').should('be.visible')
            cy.get('[data-cy="breadcrumb-country"]').should('be.visible')
            
            // Test country breadcrumb link
            cy.get('[data-cy="breadcrumb-country"] a').click()
            cy.url().should('include', '/country/')
        })

        it('should have working global navigation link', () => {
            cy.visit('/highwaytype/interstate-highway/')
            
            // Test Recent link from global nav
            cy.contains('nav a', 'Recent').click()
            cy.url().should('include', '/recent')
        })

        it('should display list of highways with shields', () => {
            cy.get('[data-cy="highway-item"]').should('be.visible')
            cy.get('[data-cy="highway-item"]').should('have.length.at.least', 1)
            cy.get('[data-cy="highway-shield"]').should('be.visible')
        })

        it('should display sign counts for highways', () => {
            cy.get('[data-cy="highway-item"]').first().should('contain', 'Sign')
        })

        it('should navigate to highway page from highway link', () => {
            cy.get('[data-cy="highway-link"]').first().click()
            cy.url().should('include', '/highway/')
        })

        it('should navigate to highway page from highway name link', () => {
            cy.get('[data-cy="highway-item"] a.hover\\:underline').first().click()
            cy.url().should('include', '/highway/')
        })
    })

    context('Highway type page - Alberta Provincial Highway', () => {
        beforeEach(() => {
            cy.visit('/highwaytype/alberta-provincial-highway/')
        })

        it('should display highway type title', () => {
            cy.get('[data-cy="highway-type-title"]').should('be.visible')
            cy.get('[data-cy="highway-type-title"]').should('contain', 'Alberta')
        })

        it('should have list of highways', () => {
            cy.get('[data-cy="highway-item"]').should('have.length.at.least', 1)
        })

        it('should have working breadcrumb to country', () => {
            cy.get('[data-cy="breadcrumb-country"]').should('be.visible')
            cy.get('[data-cy="breadcrumb-country"]').should('contain', 'Canada')
        })
    })

    context('Highway type with state groupings', () => {
        it('should display highways organized by state if multi-state', () => {
            cy.visit('/highwaytype/interstate-highway/')
            
            // Multi-state highway types may have state groupings
            cy.get('body').should('be.visible')
            cy.get('[data-cy="highway-item"]').should('have.length.at.least', 5)
        })
    })

    context('Highway type navigation from homepage', () => {
        it('should navigate to highway type from homepage', () => {
            cy.visit('/')
            
            // Click on a highway type link if present
            cy.get('[data-cy="highway-link-interstate-highway"]').click()
            cy.url().should('include', '/highwaytype/interstate-highway')
        })
    })

    context('Highway type page pagination', () => {
        it('should handle pagination if many highways in type', () => {
            cy.visit('/highwaytype/interstate-highway/')
            
            cy.get('body').then($body => {
                if ($body.find('[data-cy="pagination"]').length > 0) {
                    cy.get('[data-cy="pagination"]').should('be.visible')
                }
            })
        })
    })
})
