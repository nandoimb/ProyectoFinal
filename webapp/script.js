document.addEventListener('DOMContentLoaded', () => {
    fetchProducts();
    fetchOrders();

    const createForm = document.getElementById('product-form');
    createForm.addEventListener('submit', (event) => {
        event.preventDefault();
        createProduct();
    });

    const updateForm = document.getElementById('update-form');
    updateForm.addEventListener('submit', (event) => {
        event.preventDefault();
        updateProduct();
    });

    const deleteForm = document.getElementById('delete-form');
    deleteForm.addEventListener('submit', (event) => {
        event.preventDefault();
        deleteProduct();
    });

    const orderForm = document.getElementById('order-form');
    orderForm.addEventListener('submit', (event) => {
        event.preventDefault();
        createOrder();
    });

    const orderUpdateForm = document.getElementById('order-update-form');
    orderUpdateForm.addEventListener('submit', (event) => {
        event.preventDefault();
        updateOrder();
    });

    const orderDeleteForm = document.getElementById('order-delete-form');
    orderDeleteForm.addEventListener('submit', (event) => {
        event.preventDefault();
        deleteOrder();
    });
});

async function fetchProducts() {
    try {
        const response = await fetch('http://localhost:8080/products');
        const products = await response.json();
        displayProducts(products);
    } catch (error) {
        console.error('Error fetching products:', error);
    }
}

function displayProducts(products) {
    const productList = document.getElementById('product-list');
    productList.innerHTML = '';
    products.forEach(product => {
        const productDiv = document.createElement('div');
        productDiv.className = 'product';
        productDiv.textContent = `ID: ${product.id}, Name: ${product.name}, Price: $${product.price}`;
        productList.appendChild(productDiv);
    });
}

async function createProduct() {
    const name = document.getElementById('name').value;
    const price = document.getElementById('price').value;

    const product = {
        name,
        price: parseFloat(price)
    };

    try {
        const response = await fetch('http://localhost:8080/products', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(product)
        });

        if (response.ok) {
            fetchProducts();
            document.getElementById('product-form').reset();
        } else {
            const error = await response.json();
            console.error('Error creating product:', error);
        }
    } catch (error) {
        console.error('Error creating product:', error);
    }
}

async function updateProduct() {
    const id = document.getElementById('update-id').value;
    const name = document.getElementById('update-name').value;
    const price = document.getElementById('update-price').value;

    const product = {
        id: parseInt(id),
        name,
        price: parseFloat(price)
    };

    try {
        const response = await fetch(`http://localhost:8080/products/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(product)
        });

        if (response.ok) {
            fetchProducts();
            document.getElementById('update-form').reset();
        } else {
            const error = await response.json();
            console.error('Error updating product:', error);
        }
    } catch (error) {
        console.error('Error updating product:', error);
    }
}

async function deleteProduct() {
    const id = document.getElementById('delete-id').value;

    try {
        const response = await fetch(`http://localhost:8080/products/${id}`, {
            method: 'DELETE'
        });

        if (response.ok) {
            fetchProducts();
            document.getElementById('delete-form').reset();
        } else {
            const error = await response.json();
            console.error('Error deleting product:', error);
        }
    } catch (error) {
        console.error('Error deleting product:', error);
    }
}

async function fetchOrders() {
    try {
        const response = await fetch('http://localhost:8080/orders');
        const orders = await response.json();
        displayOrders(orders);
    } catch (error) {
        console.error('Error fetching orders:', error);
    }
}

function displayOrders(orders) {
    const orderList = document.getElementById('order-list');
    orderList.innerHTML = '';
    orders.forEach(order => {
        const orderDiv = document.createElement('div');
        orderDiv.className = 'order';
        orderDiv.textContent = `ID: ${order.id}, User ID: ${order.userID}`;
        orderList.appendChild(orderDiv);
    });
}

async function createOrder() {
    const userID = document.getElementById('user-id').value;

    const order = {
        userID: parseInt(userID)
    };

    try {
        const response = await fetch('http://localhost:8080/orders', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(order)
        });

        if (response.ok) {
            fetchOrders();
            document.getElementById('order-form').reset();
        } else {
            const error = await response.json();
            console.error('Error creating order:', error);
        }
    } catch (error) {
        console.error('Error creating order:', error);
    }
}

async function updateOrder() {
    const id = document.getElementById('order-update-id').value;
    const userID = document.getElementById('order-update-user-id').value;

    const order = {
        id: parseInt(id),
        userID: parseInt(userID)
    };

    try {
        const response = await fetch(`http://localhost:8080/orders/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(order)
        });

        if (response.ok) {
            fetchOrders();
            document.getElementById('order-update-form').reset();
        } else {
            const error = await response.json();
            console.error('Error updating order:', error);
        }
    } catch (error) {
        console.error('Error updating order:', error);
    }
}

async function deleteOrder() {
    const id = document.getElementById('order-delete-id').value;

    try {
        const response = await fetch(`http://localhost:8080/orders/${id}`, {
            method: 'DELETE'
        });

        if (response.ok) {
            fetchOrders();
            document.getElementById('order-delete-form').reset();
        } else {
            const error = await response.json();
            console.error('Error deleting order:', error);
        }
    } catch (error) {
        console.error('Error deleting order:', error);
    }
}
