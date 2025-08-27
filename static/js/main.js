class PhoneTrackingDashboard {
    constructor() {
        this.map = null;
        this.charts = {};
        this.devices = [];
        this.notifications = [];
        this.init();
    }

    init() {
        this.setupSidebar();
        this.setupModals();
        this.setupMap();
        this.setupCharts();
        this.setupRealTimeUpdates();
        this.setupSearch();
        this.loadInitialData();
    }

    setupSidebar() {
        const sidebarToggle = document.getElementById('sidebarToggle');
        const sidebar = document.getElementById('sidebar');
        const mainContent = document.getElementById('mainContent');

        sidebarToggle.addEventListener('click', () => {
            sidebar.classList.toggle('show');
            mainContent.classList.toggle('full-width');
        });

        // Close sidebar on mobile when clicking outside
        document.addEventListener('click', (e) => {
            if (window.innerWidth <= 992) {
                if (!sidebar.contains(e.target) && !sidebarToggle.contains(e.target) && sidebar.classList.contains('show')) {
                    sidebar.classList.remove('show');
                    mainContent.classList.remove('full-width');
                }
            }
        });

        // Handle nav link clicks
        document.querySelectorAll('.nav-link').forEach(link => {
            link.addEventListener('click', (e) => {
                e.preventDefault();
                document.querySelectorAll('.nav-link').forEach(l => l.classList.remove('active'));
                link.classList.add('active');
                this.showToast('Navigation updated', 'info');

                // Close sidebar on mobile after selection
                if (window.innerWidth <= 992) {
                    sidebar.classList.remove('show');
                    mainContent.classList.remove('full-width');
                }
            });
        });
    }

    setupModals() {
        const addDeviceBtn = document.getElementById('addDeviceBtn');
        const addDeviceModal = document.getElementById('addDeviceModal');
        const closeModal = document.getElementById('closeModal');
        const cancelAdd = document.getElementById('cancelAdd');
        const settingsBtn = document.getElementById('settingsBtn');
        const settingsModal = document.getElementById('settingsModal');
        const closeSettingsModal = document.getElementById('closeSettingsModal');
        const cancelSettings = document.getElementById('cancelSettings');

        // Add device modal
        addDeviceBtn.addEventListener('click', () => {
            addDeviceModal.classList.add('active');
        });

        [closeModal, cancelAdd].forEach(btn => {
            btn.addEventListener('click', () => {
                addDeviceModal.classList.remove('active');
            });
        });

        // Settings modal
        settingsBtn.addEventListener('click', (e) => {
            e.preventDefault();
            settingsModal.classList.add('active');
        });

        [closeSettingsModal, cancelSettings].forEach(btn => {
            btn.addEventListener('click', () => {
                settingsModal.classList.remove('active');
            });
        });

        // Close modals when clicking outside
        document.querySelectorAll('.modal-overlay').forEach(overlay => {
            overlay.addEventListener('click', (e) => {
                if (e.target === overlay) {
                    overlay.classList.remove('active');
                }
            });
        });

        // Tab functionality
        document.querySelectorAll('.tab-btn').forEach(btn => {
            btn.addEventListener('click', () => {
                const tabId = btn.dataset.tab;
                document.querySelectorAll('.tab-btn').forEach(b => b.classList.remove('active'));
                document.querySelectorAll('.tab-content').forEach(c => c.style.display = 'none');
                btn.classList.add('active');
                document.getElementById(tabId).style.display = 'block';
            });
        });

        // Form submissions
        document.getElementById('addDeviceForm').addEventListener('submit', (e) => {
            e.preventDefault();
            this.addDevice(new FormData(e.target));
            addDeviceModal.classList.remove('active');
        });
    }

    setupMap() {
        this.map = L.map('map').setView([27.7172, 85.3240], 13); // Kathmandu coordinates

        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '© OpenStreetMap contributors'
        }).addTo(this.map);

        // Add device markers
        const devices = [
            { name: "John's iPhone", lat: 27.7172, lng: 85.3240, status: 'online', battery: 85 },
            { name: "Sarah's Galaxy", lat: 27.6588, lng: 85.3247, status: 'online', battery: 92 },
            { name: "Mike's Pixel", lat: 27.6710, lng: 85.4298, status: 'offline', battery: 15 },
            { name: "Lisa's OnePlus", lat: 28.2096, lng: 83.9856, status: 'online', battery: 67 }
        ];

        devices.forEach(device => {
            const color = device.status === 'online' ? '#4ecdc4' : '#ff6b6b';
            const marker = L.circleMarker([device.lat, device.lng], {
                radius: 8,
                fillColor: color,
                color: '#fff',
                weight: 2,
                opacity: 1,
                fillOpacity: 0.8
            }).addTo(this.map);

            marker.bindPopup(`
                <div style="text-align: center;">
                    <strong>${device.name}</strong><br>
                    Status: <span style="color: ${color}">${device.status}</span><br>
                    Battery: ${device.battery}%
                </div>
            `);
        });
    }

    setupCharts() {
        // Activity Chart
        const activityCtx = document.getElementById('activityChart').getContext('2d');
        this.charts.activity = new Chart(activityCtx, {
            type: 'line',
            data: {
                labels: Array.from({ length: 24 }, (_, i) => `${i}:00`),
                datasets: [{
                    label: 'Active Devices',
                    data: [8, 9, 7, 10, 12, 11, 9, 8, 10, 12, 14, 13, 12, 11, 10, 12, 11, 9, 8, 7, 9, 10, 8, 7],
                    borderColor: '#4a6cf7',
                    backgroundColor: 'rgba(74, 108, 247, 0.1)',
                    borderWidth: 3,
                    fill: true,
                    tension: 0.4
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    legend: {
                        display: false
                    }
                },
                scales: {
                    y: {
                        beginAtZero: true,
                        grid: {
                            color: 'rgba(0, 0, 0, 0.05)'
                        }
                    },
                    x: {
                        grid: {
                            display: false
                        }
                    }
                }
            }
        });

        // Battery Chart
        const batteryCtx = document.getElementById('batteryChart').getContext('2d');
        this.charts.battery = new Chart(batteryCtx, {
            type: 'doughnut',
            data: {
                labels: ['85-100%', '70-84%', '50-69%', '20-49%', '0-19%'],
                datasets: [{
                    data: [3, 4, 3, 1, 1],
                    backgroundColor: [
                        '#4a6cf7',
                        '#17a2b8',
                        '#ffc107',
                        '#fd7e14',
                        '#dc3545'
                    ],
                    borderWidth: 0
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    legend: {
                        position: 'bottom',
                        labels: {
                            padding: 20,
                            usePointStyle: true,
                            pointStyle: 'circle'
                        }
                    }
                }
            }
        });
    }

    setupRealTimeUpdates() {
        // Simulate real-time updates
        setInterval(() => {
            this.updateDeviceStats();
            this.updateMap();
        }, 10000);

        // Simulate new notifications
        setInterval(() => {
            this.addRandomNotification();
        }, 30000);
    }

    setupSearch() {
        const searchInput = document.querySelector('.search-input');
        searchInput.addEventListener('input', (e) => {
            const query = e.target.value.toLowerCase();
            this.filterDevices(query);
        });
    }

    loadInitialData() {
        // Simulate loading initial data
        setTimeout(() => {
            this.showToast('Dashboard loaded successfully', 'success');
        }, 1000);
    }

    updateDeviceStats() {
        // Update active devices count
        const activeDevicesEl = document.querySelector('.stat-value');
        const currentCount = parseInt(activeDevicesEl.textContent);
        const newCount = Math.max(8, Math.min(15, currentCount + Math.floor(Math.random() * 3) - 1));
        activeDevicesEl.textContent = newCount;

        // Update charts with new data
        if (this.charts.activity) {
            const newData = this.charts.activity.data.datasets[0].data.slice(1);
            newData.push(Math.floor(Math.random() * 6) + 8);
            this.charts.activity.data.datasets[0].data = newData;
            this.charts.activity.update('none');
        }
    }

    updateMap() {
        // Simulate device movement (in a real app, this would come from API)
        console.log('Updating device positions...');
    }

    addRandomNotification() {
        const notifications = [
            { type: 'warning', icon: 'battery-empty', title: 'Low Battery', message: 'Device battery below 20%' },
            { type: 'info', icon: 'map-pin', title: 'Location Update', message: 'Device location updated' },
            { type: 'error', icon: 'wifi', title: 'Connection Lost', message: 'Device went offline' }
        ];

        const notification = notifications[Math.floor(Math.random() * notifications.length)];
        this.showToast(notification.message, notification.type);
    }

    addDevice(formData) {
        const deviceData = Object.fromEntries(formData);
        console.log('Adding device:', deviceData);
        this.showToast('Device added successfully', 'success');
    }

    filterDevices(query) {
        const deviceItems = document.querySelectorAll('.device-item');
        deviceItems.forEach(item => {
            const text = item.textContent.toLowerCase();
            item.style.display = text.includes(query) ? 'flex' : 'none';
        });
    }

    showToast(message, type = 'info') {
        const toastContainer = document.getElementById('toastContainer');
        const toast = document.createElement('div');
        toast.className = `toast ${type}`;

        const icons = {
            success: 'check-circle',
            error: 'exclamation-circle',
            warning: 'exclamation-triangle',
            info: 'info-circle'
        };

        toast.innerHTML = `
            <i class="fas fa-${icons[type]} toast-icon"></i>
            <div class="toast-content">
                <div class="toast-title">${type.charAt(0).toUpperCase() + type.slice(1)}</div>
                <div class="toast-message">${message}</div>
            </div>
            <button class="toast-close">&times;</button>
        `;

        toastContainer.appendChild(toast);

        // Show toast with animation
        setTimeout(() => {
            toast.classList.add('show');
        }, 10);

        // Add close button functionality
        toast.querySelector('.toast-close').addEventListener('click', () => {
            toast.classList.remove('show');
            setTimeout(() => {
                toast.remove();
            }, 300);
        });

        // Auto remove after 5 seconds
        setTimeout(() => {
            if (toast.parentNode) {
                toast.classList.remove('show');
                setTimeout(() => {
                    toast.remove();
                }, 300);
            }
        }, 5000);
    }
}

// Initialize dashboard when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    new PhoneTrackingDashboard();
});