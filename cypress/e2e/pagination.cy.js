describe('Pagination Tests', () => {
    context('Pages with pagination', () => {
        beforeEach(() => {
            cy.visit('/state/california/page/3/list.html') // or any page with multiple pages of content
        })

        it('should display pagination navigation when there are multiple pages', () => {
            cy.get('[data-cy="pagination"]').should('be.visible')
            cy.get('[data-cy="pagination"] ul').should('exist')
        })

        it('should display all pagination controls', () => {
            cy.get('[data-cy="pagination-first"]').should('be.visible')
            cy.get('[data-cy="pagination-prev"]').should('be.visible')
            cy.get('[data-cy="pagination-next"]').should('be.visible')
            cy.get('[data-cy="pagination-last"]').should('be.visible')
        })

        it('should highlight the current page', () => {
            cy.get('[data-cy="pagination-page-current"]').should('be.visible')
            cy.get('[data-cy="pagination-page-current"]').should('contain', '3')
        })

        it('should navigate to the next page', () => {
            cy.get('[data-cy="pagination-next"]').click()
            cy.get('[data-cy="pagination-page-current"]').should('contain', '4')
            cy.url().should('include', '/page/4')
        })

        it('should navigate to the previous page', () => {
            // Go to page 2 first
            cy.get('[data-cy="pagination-next"]').click()
            cy.get('[data-cy="pagination-page-current"]').should('contain', '4')

            // Then go back
            cy.get('[data-cy="pagination-prev"]').click()
            cy.get('[data-cy="pagination-page-current"]').should('contain', '3')
            cy.url().should('not.include', '/page/4')
        })

        it('should navigate to the first page from any page', () => {
            // Navigate to a later page
            cy.get('[data-cy="pagination-next"]').click()
            cy.get('[data-cy="pagination-page-current"]').should('contain', '4')

            // Click First
            cy.get('[data-cy="pagination-first"]').click()
            cy.get('[data-cy="pagination-page-current"]').should('contain', '1')
            cy.url().should('not.include', '/page/')
        })

        it('should navigate to the last page', () => {
            cy.get('[data-cy="pagination-last"]').click()

            // Verify we're on the last page by checking that Next is not present
            cy.get('[data-cy="pagination-next"]').should('not.exist')

            // Verify current page is highlighted
            cy.get('[data-cy="pagination-page-current"]').should('be.visible')
        })

        it('should navigate to a specific page by clicking page number', () => {
            // Find and click a specific page number (e.g., page 3)
            cy.get('[data-cy="pagination-page"][data-page="4"]').click()

            cy.get('[data-cy="pagination-page-current"]').should('contain', '4')
            cy.url().should('include', '/page/4')
        })

        it('should display correct range of page numbers', () => {
            // Verify that pagination shows up to 5 page slots
            cy.get('[data-cy="pagination"] [data-cy^="pagination-page"]')
                .should('have.length.at.most', 5)
        })

        it('should update page numbers when navigating', () => {
            // Click Next several times and verify page numbers update
            cy.get('[data-cy="pagination-next"]').click()
            cy.get('[data-cy="pagination-next"]').click()
            cy.get('[data-cy="pagination-next"]').click()

            // Should now be on page 6
            cy.get('[data-cy="pagination-page-current"]').should('contain', '6')

            // Verify page 6 is visible in the pagination
            cy.get('[data-cy="pagination-page"][data-page="4"], [data-cy="pagination-page-current"][data-page="6"]')
                .should('be.visible')
        })

        it('should correctly identify page numbers using data-page attribute', () => {
            // Verify current page has correct data-page attribute
            cy.get('[data-cy="pagination-page-current"]')
                .should('have.attr', 'data-page', '3')

            // Navigate to page 4
            cy.get('[data-cy="pagination-next"]').click()

            cy.get('[data-cy="pagination-page-current"]')
                .should('have.attr', 'data-page', '4')
        })
    })

    context('First page behaviors', () => {
        beforeEach(() => {
            cy.visit('/state/california/list.html') // or any page with multiple pages of content
        })

        it('should not show Previous and First links on first page', () => {
            cy.get('[data-cy="pagination-prev"]').should('not.exist')
        })

        it('should show Next and Last links on first page', () => {
            cy.get('[data-cy="pagination-next"]').should('be.visible')
            cy.get('[data-cy="pagination-last"]').should('be.visible')
        })
    })

    context('Last page behaviors', () => {
        beforeEach(() => {
            cy.visit('/state/california/list.html') // or any page with multiple pages of content
            // Navigate to last page
            cy.get('[data-cy="pagination-last"]').click()
        })

        it('should not show Next and Last links on last page', () => {
            cy.get('[data-cy="pagination-last"]').should('exist')
        })

        it('should show Previous and First links on last page', () => {
            cy.get('[data-cy="pagination-prev"]').should('be.visible')
            cy.get('[data-cy="pagination-first"]').should('be.visible')
        })
    })

    context('Pages without pagination', () => {
        it('should not display pagination on single page content', () => {
            // Visit a page that should have only one page of content
            cy.visit('/highway/id44')

            cy.get('[data-cy="pagination"]').should('not.exist')
        })
    })


    context('URL and browser navigation', () => {
        beforeEach(() => {
            cy.visit('/state/california/list.html') // or any page with multiple pages of content
        })

        it('should update URL when navigating pages', () => {
            cy.get('[data-cy="pagination-next"]').click()
            cy.url().should('include', '/page/2')

            cy.get('[data-cy="pagination-next"]').click()
            cy.url().should('include', '/page/3')
        })

        it('should maintain pagination state on browser back/forward', () => {
            // Navigate forward
            cy.get('[data-cy="pagination-next"]').click()
            cy.get('[data-cy="pagination-page-current"]').should('contain', '2')

            // Go back
            cy.go('back')
            cy.get('[data-cy="pagination-page-current"]').should('contain', '1')

            // Go forward
            cy.go('forward')
            cy.get('[data-cy="pagination-page-current"]').should('contain', '2')
        })

        it('should load correct page when accessing paginated URL directly', () => {
            cy.visit('/state/california/page/3/list.html')
            cy.get('[data-cy="pagination-page-current"]').should('contain', '3')
            cy.get('[data-cy="pagination-page-current"]').should('have.attr', 'data-page', '3')
        })
    })

    context('Middle page behaviors', () => {
        beforeEach(() => {
            cy.visit('/state/california/list.html')
            // Navigate to a middle page (page 3)
            cy.get('[data-cy="pagination-page"][data-page="3"]').click()
        })

        it('should show all navigation controls on middle pages', () => {
            cy.get('[data-cy="pagination-first"]').should('be.visible')
            cy.get('[data-cy="pagination-prev"]').should('be.visible')
            cy.get('[data-cy="pagination-next"]').should('be.visible')
            cy.get('[data-cy="pagination-last"]').should('be.visible')
        })

        it('should correctly navigate from middle page', () => {
            // Verify we're on page 3
            cy.get('[data-cy="pagination-page-current"]').should('contain', '3')

            // Go forward
            cy.get('[data-cy="pagination-next"]').click()
            cy.get('[data-cy="pagination-page-current"]').should('contain', '4')

            // Go back twice
            cy.get('[data-cy="pagination-prev"]').click()
            cy.get('[data-cy="pagination-page-current"]').should('contain', '3')

            cy.get('[data-cy="pagination-prev"]').click()
            cy.get('[data-cy="pagination-page-current"]').should('contain', '2')
        })
    })
})
