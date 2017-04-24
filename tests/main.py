import requests
import unittest
import time

BASE_URL = 'http://127.0.0.1:1234'

TASK = '{"name": "TEST", "completed": false, "due": "2000-01-01T00:00:00.000000Z"}'
TASK_UPDATE = '{"name": "TEST", "completed": true, "due": "2000-01-01T00:00:00.000000Z"}'

class TestTaskAPI(unittest.TestCase):
    def setUp(self):

        response = requests.post('%s/%s' % (BASE_URL, 'tasks'), data=TASK)
        self.assertEqual(response.status_code, 201)
        self.id = response.json()['id']

    def test_isalive(self):
        response = requests.get(BASE_URL)
        self.assertEqual(response.status_code, 200)

    def test_list(self):
        response = requests.get('%s/%s' % (BASE_URL, 'tasks'))
        self.assertEqual(response.status_code, 200)

        for t in response.json():
            if t['id'] == self.id:
                self.assertEqual(t['name'], 'TEST')

    def test_read(self):
        response = requests.get('%s/%s/%d' % (BASE_URL, 'tasks', self.id))
        self.assertEqual(response.status_code, 200)
        self.assertEqual(response.json()['name'], 'TEST')

    def test_update(self):
        response = requests.put('%s/%s/%d' % (BASE_URL, 'tasks', self.id), data=TASK_UPDATE)
        self.assertEqual(response.status_code, 201)
        self.assertEqual(response.json()['completed'], True)

    def tearDown(self):
        response = requests.delete('%s/%s/%d' % (BASE_URL, 'tasks', self.id))
        self.assertEqual(response.status_code, 200)

        response = requests.get('%s/%s/%d' % (BASE_URL, 'tasks', self.id))
        self.assertEqual(response.status_code, 404)

if __name__ == '__main__':
    unittest.main()
