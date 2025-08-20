describe('Sign Page Tests', () => {
    // Visit a specific sign page before each test
    beforeEach(() => {
        // Assuming a specific sign URL structure
        cy.visit('/sign/909981781')
    })

    it('should display the sign title and image', () => {
        // Test basic elements that should be on every sign page
        cy.get('[data-cy="sign-title"]').should('be.visible')
        cy.get('[data-cy="sign-image"]').should('be.visible')
        cy.get('[data-cy="sign-image"]').should('have.attr', 'alt')
    })

    it('should display metadata about the sign', () => {
        cy.get('[data-cy="sign-metadata"]').should('be.visible')
        cy.get('[data-cy="sign-date-taken"]').should('exist')
        cy.get('[data-cy="sign-location"]').should('exist')
    })

    it('should have working highway links', () => {
        cy.get('[data-cy="highway-links"]').should('be.visible')
        // Get the first highway link and click it
        cy.get('[data-cy="highway-link"]').first().click()
        // URL should change to a highway page
        cy.url().should('include', '/highway/')
    })

    it('should have a working state link', () => {
        cy.get('[data-cy="state-link"]').click()
        cy.url().should('include', '/state/')
    })

    it('should navigate between related signs if available', () => {
        // Test navigation to other signs if they exist
        cy.get('[data-cy="related-signs"]').then($relatedSigns => {
            if ($relatedSigns.find('[data-cy="next-sign"]').length > 0) {
                cy.get('[data-cy="next-sign"]').click()
                cy.url().should('include', '/sign/')
                cy.url().should('not.include', '/sign/i90-exit34')
            }
        })
    })

    it('should display features associated with the sign', () => {
        cy.get('[data-cy="features-section"]').then($featuresSection => {
            if ($featuresSection.length > 0) {
                cy.get('[data-cy="feature-link"]').should('exist')
                cy.get('[data-cy="feature-link"]').first().click()
                cy.url().should('include', '/feature/')
            }
        })
    })
})

describe('Sign Navigation Between Pages', () => {
    it('should navigate to a sign page from the homepage', () => {
        cy.visit('/')
        // Assuming you have recent or featured signs on homepage
        cy.get('[data-cy="featured-sign"]').first().click()
        cy.url().should('include', '/sign/')
        cy.get('[data-cy="sign-title"]').should('be.visible')
    })

    it('should navigate to a sign page from a highway page', () => {
        cy.visit('/highway/i90')
        cy.get('[data-cy="highway-sign-item"]').first().click()
        cy.url().should('include', '/sign/')
        cy.get('[data-cy="sign-title"]').should('be.visible')
    })
})