from app.database.connection import get_connection

def get_personalized_recommendations(user_id):
    conn = get_connection()
    cursor = conn.cursor()
    cursor.execute("SELECT product_id, product_name, image_url FROM recommendations WHERE user_id = %s", (user_id,))
    rows = cursor.fetchall()
    conn.close()
    return [{"product_id": r[0], "product_name": r[1], "image_url": r[2]} for r in rows]

def get_related_products(product_id):
    conn = get_connection()
    cursor = conn.cursor()
    cursor.execute("SELECT product_id, product_name, image_url FROM related_products WHERE reference_product_id = %s", (product_id,))
    rows = cursor.fetchall()
    conn.close()
    return [{"product_id": r[0], "product_name": r[1], "image_url": r[2]} for r in rows]

def save_feedback(user_id, product_id, feedback_type):
    conn = get_connection()
    cursor = conn.cursor()
    cursor.execute(
        "INSERT INTO feedback (user_id, product_id, feedback_type) VALUES (%s, %s, %s)",
        (user_id, product_id, feedback_type)
    )
    conn.commit()
    conn.close()
