from fastapi import APIRouter

from app.controllers.get_all_categories import router as get_all
from app.controllers.get_category_by_id import router as get_by_id
from app.controllers.create_category import router as create
from app.controllers.update_category import router as update
from app.controllers.delete_category import router as delete
from app.controllers.get_category_products import router as get_products

router = APIRouter()
router.include_router(get_all)
router.include_router(get_by_id)
router.include_router(create)
router.include_router(update)
router.include_router(delete)
router.include_router(get_products)
